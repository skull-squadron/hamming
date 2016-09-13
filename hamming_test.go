//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 Barry Allard
//
// MIT license
//
package hamming

import (
	"math"
	"strconv"
	"testing"
)

const (
	maxInt  = (1 << (strconv.IntSize - 1)) - 1
	maxUint = (1 << strconv.IntSize) - 1
)

func TestCountBitInt8(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt8 {
			continue
		}
		if actualN := CountBitsInt8(int8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitInt16(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt16 {
			continue
		}
		if actualN := CountBitsInt16(int16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitInt32(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := CountBitsInt32(int32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitInt64(t *testing.T) {
	for _, c := range testCountBitsCases {
		if actualN := CountBitsInt64(int64(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitInt(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > maxInt {
			continue
		}
		if actualN := CountBitsInt(int(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitUint8(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := CountBitsUint8(uint8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitUint16(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint16 {
			continue
		}
		if actualN := CountBitsUint16(uint16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitUint32(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint32 {
			continue
		}
		if actualN := CountBitsUint32(uint32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitUint64(t *testing.T) {
	for _, c := range testCountBitsCases {
		if actualN := CountBitsUint64(c.x); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitUint(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > maxUint {
			continue
		}
		if actualN := CountBitsUint(uint(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitByte(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := CountBitsByte(byte(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitByteAlt(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := CountBitsByteAlt(byte(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitRune(t *testing.T) {
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := CountBitsRune(rune(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

// ============== benchmarks ==============

func BenchmarkCountBitsInt8(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsInt8(int8(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsInt16(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsInt16(int16(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsInt32(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsInt32(int32(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsInt64(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsInt64(int64(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsInt(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsInt(int(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsUint16(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsUint16(uint16(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsUint32(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsUint32(uint32(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsUint64(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsUint64(uint64(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsUint64Alt(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsUint64Alt(uint64(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsUint(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsUint(uint(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsByte(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsByte(byte(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsByteAlt(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsByteAlt(byte(i))
	}

	writeToTempFile(z)
}

func BenchmarkCountBitsRune(b *testing.B) {
	z := 0

	for i := 0; i < b.N; i++ {
		z += CountBitsRune(rune(i))
	}

	writeToTempFile(z)
}
