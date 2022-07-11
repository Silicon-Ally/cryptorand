package cryptorand

import "testing"

func TestSourceInt63(t *testing.T) {
	testSource(t, func(src *Source) uint64 { return uint64(src.Int63()) }, 0x7FFFFFFFFFFFFFFF)
}

func TestSourceUint64(t *testing.T) {
	testSource(t, (*Source).Uint64, 0xFFFFFFFFFFFFFFFF)
}

func testSource(t *testing.T, fn func(*Source) uint64, bitOrCheckTarget uint64) {
	// There's not really a great way to test random number generation, so we
	// just generate ten thousand numbers and do some basic checks:
	//  1. That there are no collisions, and
	//  2. That at least one of the values sets a bit in each target position
	// Either of these conditions not being satisfied is overwhelmingly likely to
	// indicate a bug, neither test should fail due to flakiness before the heat
	// death of the universe. If it does, we should play the lottery and/or
	// examine our beliefs around higher powers.

	src := NewSource()
	freq := make(map[uint64]int)
	for i := 0; i < 10000; i++ {
		freq[fn(src)]++
	}

	var bitOrCheck uint64
	bitAndCheck := bitOrCheckTarget
	for val, cnt := range freq {
		if cnt != 1 {
			t.Fatalf("value %d appeared %d times", val, cnt)
		}

		bitOrCheck |= val
		bitAndCheck &= val
	}

	if got, want := bitOrCheck, bitOrCheckTarget; got != want {
		// The odds of all 10,000 runs returning '0' for a position are so
		// improbable that it is guaranteed to be a bug if it happens.
		t.Fatalf("at least one target bit wasn't set among random values, got %d, want %d", got, want)
	}
	if bitAndCheck != 0 {
		t.Fatalf("at least one target bit was set for all values, got %d", bitAndCheck)
	}
}

func TestRand(t *testing.T) {
	r := New()

	n := 10
	freq := make(map[int]int)
	for i := 0; i < 10000; i++ {
		freq[r.Intn(n)]++
	}

	for i := 0; i < n; i++ {
		if freq[i] == 0 {
			t.Fatalf("didn't generate value %d in 10,000 random runs", i)
		}
	}
}
