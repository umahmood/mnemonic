package mnemonic_test

import (
	"testing"

	"github.com/umahmood/mnemonic"
)

func TestNewWithValidConfig(t *testing.T) {
	_, err := mnemonic.New(mnemonic.DefaultConfig)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestNewWithInvalidConfig(t *testing.T) {
	_, err := mnemonic.New(mnemonic.Config{
		Bits: 0,
	})
	if err != mnemonic.ErrInvalidBitSize {
		t.Errorf("%v", err)
	}
	_, err = mnemonic.New(mnemonic.Config{
		Bits: 42,
	})
	if err != mnemonic.ErrInvalidBitSize {
		t.Errorf("%v", err)
	}
}

func TestWordsWithMultiplesOf32(t *testing.T) {
	bitsToWords := map[int]int{
		32:   3,
		64:   6,
		96:   9,
		128:  12,
		256:  24,
		1568: 147,
	}
	for bits, numWords := range bitsToWords {
		m, err := mnemonic.New(mnemonic.Config{
			Bits: bits,
		})
		if err != nil {
			t.Errorf("%v", err)
		}
		words, err := m.Words()
		if err != nil {
			t.Errorf("%v", err)
		}
		if len(words) != numWords {
			t.Errorf("fail: mismatch between config. bits %d and words %d", bits, numWords)
		}
	}
}
