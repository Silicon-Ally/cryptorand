// Package cryptorand provides a *rand.Rand implementation backed by the
// standard library's 'crypto/rand' package.
package cryptorand

import (
	"encoding/binary"
	"math/rand"

	cryptorand "crypto/rand"
)

type Source struct{}

// New returns a *rand.Rand backed by a secure random number generator.
func New() *rand.Rand {
	return rand.New(&Source{})
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() & 0x7FFFFFFFFFFFFFFF)
}

func (*Source) Uint64() uint64 {
	var buf [8]byte
	n, err := cryptorand.Read(buf[:])
	if err != nil {
		panic(err)
	}
	if n != 8 {
		panic("read incorrect number of bytes")
	}
	return binary.BigEndian.Uint64(buf[:])
}

func (*Source) Seed(_ int64) {}
