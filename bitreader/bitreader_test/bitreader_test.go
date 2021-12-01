package bitreader_test

import (
	"bytes"
	"io"
	"testing"

	bitreader "github.com/umahmood/mnemonic/bitreader"
)

func TestReadNegativeBits(t *testing.T) {
	var (
		input      = []byte{80, 176, 183, 84, 3}
		bitsToRead = -42
	)
	var want error
	bits := bitreader.NewBitReader(bytes.NewReader(input))
	_, err := bits.ReadBits(bitsToRead)
	if err != want {
		t.Errorf("err is not nil got %v", err)
	}
}

func TestReadUntilEOF(t *testing.T) {
	var (
		input      = []byte{99, 255}
		bitsToRead = 8
		want       = io.EOF
	)
	bits := bitreader.NewBitReader(bytes.NewReader(input))
	_, err := bits.ReadBits(bitsToRead)
	_, err = bits.ReadBits(bitsToRead)
	_, err = bits.ReadBits(bitsToRead)
	if err != want {
		t.Errorf("err is not eof got %v", err)
	}
}

func TestZeroBitsInput(t *testing.T) {
	var (
		input      = []byte{}
		bitsToRead = 10
		want       = io.EOF
	)
	bits := bitreader.NewBitReader(bytes.NewReader(input))
	_, err := bits.ReadBits(bitsToRead)
	if err != want {
		t.Errorf("err is not eof got %v", err)
	}
}

func TestInvalidBitsInput(t *testing.T) {
	var (
		input      = []byte{}
		bitsToRead = 10
		want       = io.EOF
	)
	bits := bitreader.NewBitReader(bytes.NewReader(input))
	_, err := bits.ReadBits(bitsToRead)
	if err != want {
		t.Errorf("err is not eof got %v", err)
	}
}

func TestValidBitsInput(t *testing.T) {
	var (
		input      = []byte{42, 167, 88} // 42: 00101010 | 167: 10100111 | 88: 01011000
		bitsToRead = 12
		want       = uint64(682)
	)
	bits := bitreader.NewBitReader(bytes.NewReader(input))
	n, err := bits.ReadBits(bitsToRead)
	if err != nil {
		t.Errorf("err is not nil %v", err)
	}
	if n != want {
		t.Errorf("got %d want %d", n, want)
	}
}
