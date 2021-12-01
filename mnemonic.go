package mnemonic

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"io"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	bitreader "github.com/umahmood/mnemonic/bitreader"
)

// ErrInvalidBitSize error when bits are not a multiple of 32
var ErrInvalidBitSize = errors.New("bits must be greater than 0 and multiple of 32")

// DefaultConfig to create mnemonics, ideal entropy should be 128 to 256 bits.
var DefaultConfig = Config{
	Bits:       128,
	Passphrase: "",
}

// Config instance
type Config struct {
	Bits       int    // Bits entropy to generate, must be a multiple of 32 bits.
	Passphrase string // Optional passphrase allows you to modify the final seed.
}

// Mnemonic instance
type Mnemonic struct {
	config       Config
	groupLen     int
	checksumBits int
	checksum     int
	entropySize  int
	entropy      []byte
	seed         []byte
	words        []string
}

// New mnemonic instance
func New(c Config) (*Mnemonic, error) {
	if c.Bits == 0 || c.Bits%32 != 0 {
		return nil, ErrInvalidBitSize
	}
	return &Mnemonic{
		config:       c,
		groupLen:     11,
		checksumBits: c.Bits / 32,
		entropySize:  c.Bits / 8,
		entropy:      make([]byte, c.Bits/8),
	}, nil
}

// Words returns a randomly generated sequence of (mnemonic) words
func (m *Mnemonic) Words() ([]string, error) {
	// generate entropy
	rand.Read(m.entropy)
	// add checksum to entropy
	sum := sha256.Sum256(m.entropy)
	// take 1 bit of that hash for every 32 bits of entropy
	var (
		hash       = bitreader.NewBitReader(bytes.NewReader(sum[:]))
		bitsToRead = len(m.entropy) * 8 / 32
		checksum   = make([]byte, 0)
	)
	for bitsToRead > 0 {
		var digit uint64
		if bitsToRead > 8 {
			digit, _ = hash.ReadBits(8)
		} else {
			digit, _ = hash.ReadBits(bitsToRead)
			digit = digit << uint((8 - bitsToRead))
		}
		checksum = append(checksum, byte(digit))
		bitsToRead -= 8
	}
	// add checksums to the end of our entropy
	m.entropy = append(m.entropy, checksum...)
	// split this in to groups of 11 bits
	var (
		bits        = bitreader.NewBitReader(bytes.NewReader(m.entropy))
		wordsToRead = ((len(m.entropy)*8)/m.groupLen + 1)
		indexes     = []int{}
	)
	for wordsToRead > 0 {
		idx, err := bits.ReadBits(m.groupLen)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// convert these to decimal numbers
		indexes = append(indexes, int(idx))
		// use those numbers to select the corresponding words
		m.words = append(m.words, getWord(int(idx)))
		wordsToRead--
	}
	return m.words, nil
}

// Seed creates a 64 byte seed from underlying mnemonic words. The optional
// passphrase allows you to modify the final seed. Must call Words before this
// method otherwise seed will be nil.
func (m *Mnemonic) Seed() []byte {
	if len(m.words) != 0 && len(m.seed) == 0 {
		salt := []byte("mnemonic")
		salt = append(salt, []byte(m.config.Passphrase)...)
		m.seed = pbkdf2.Key([]byte(strings.Join(m.words, " ")), salt, 2048, 64, sha512.New)
	}
	return m.seed
}
