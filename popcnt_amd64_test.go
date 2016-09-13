//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 Barry Allard
//
// MIT license
//

// +build amd64 amd64p32 !purego

package hamming

import (
	"math"
	"testing"
)

func TestPopCntInt8(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt8 {
			continue
		}
		if actualN := PopCntInt8(int8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntInt16(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt16 {
			continue
		}
		if actualN := PopCntInt16(int16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntInt32(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := PopCntInt32(int32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntInt64(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt64 {
			continue
		}
		if actualN := PopCntInt64(int64(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntInt(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > maxInt {
			continue
		}
		if actualN := PopCntInt(int(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntUint8(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := PopCntUint8(uint8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntUint16(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint16 {
			continue
		}
		if actualN := PopCntUint16(uint16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntUint32(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint32 {
			continue
		}
		if actualN := PopCntUint32(uint32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntUint64(t *testing.T) {
	for _, c := range testCountBitsCases {
		if actualN := PopCntUint64(c.x); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntByte(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := PopCntByte(byte(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestPopCntRune(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := PopCntRune(rune(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

// ============== benchmarks ==============

func BenchmarkPopCntInt8(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntInt8(int8(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntInt16(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntInt16(int16(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntInt32(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntInt32(int32(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntInt64(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntInt64(int64(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntInt(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntInt(i)
	}

	writeToTempFile(z)
}

func BenchmarkPopCntUint8(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntUint8(uint8(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntUint16(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntUint16(uint16(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntUint32(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntUint32(uint32(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntUint64(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntUint64(uint64(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntUint(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntUint(uint(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntByte(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntByte(byte(i))
	}

	writeToTempFile(z)
}

func BenchmarkPopCntRune(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += PopCntRune(rune(i))
	}

	writeToTempFile(z)
}
