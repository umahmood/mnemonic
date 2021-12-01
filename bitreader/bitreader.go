package mnemonic

import "io"

// Bit represents one or a zero
type Bit bool

const (
	Zero Bit = false
	One  Bit = true
)

// BitReader instance
type BitReader struct {
	r io.Reader
	b [1]byte
	c uint8
}

// NewBitReader returns a new BitReader instance
func NewBitReader(r io.Reader) *BitReader {
	return &BitReader{
		r: r,
	}
}

// readbit a single bit
func (b *BitReader) readbit() (Bit, error) {
	if b.c == 0 {
		if n, err := b.r.Read(b.b[:]); n != 1 || (err != nil && err != io.EOF) {
			return Zero, err
		}
		b.c = 8
	}
	b.c--
	x := (b.b[0] & 0x80)
	b.b[0] <<= 1
	return x != 0, nil
}

// ReadBits read n bits
func (b *BitReader) ReadBits(n int) (uint64, error) {
	var u uint64
	for n > 0 {
		q, err := b.readbit()
		if err != nil {
			return 0, err
		}
		u <<= 1
		if q {
			u |= 1
		}
		n--
	}
	return u, nil
}
