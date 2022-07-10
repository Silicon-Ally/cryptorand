// Package cryptorand provides a *rand.Rand implementation backed by the
// standard library's 'crypto/rand' package.
package cryptorand

import (
	"encoding/binary"
	"math/rand"

	cryptorand "crypto/rand"
)

// Source implements math/rand.Source and math/rand.Source64, backed by the
// crypto/rand package.
type Source struct{}

// New returns a *rand.Rand backed by the crypto/rand secure random number
// generator.
func New() *rand.Rand {
	return rand.New(&Source{})
}

// NewSource returns an initialized crypto/rand-backed implementation of
// math/rand.Source.
func NewSource() *Source {
	return &Source{}
}

// Int63 returns a non-negative cryptographically secure random 63-bit integer
// as an int64.
func (s *Source) Int63() int64 {
	return int64(s.Uint64() & 0x7FFFFFFFFFFFFFFF)
}

// Uint64 returns a cruptographically secure random 64-bit value as a uint64.
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

// Seed is a no-op, as the crytographically secure random number generated
// cannot be seeded.
func (*Source) Seed(_ int64) {}
