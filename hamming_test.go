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
	"time"
)

const (
	maxInt  = (1 << (strconv.IntSize - 1)) - 1
	maxUint = (1 << strconv.IntSize) - 1
)

type testCountCase struct {
	x uint64
	n int
}

var testCountBitsCases = []testCountCase{
	{0x00, 0},
	{0x01, 1},
	{0x02, 1},
	{0x03, 2},
	{0x77, 6},
	{0xaa, 4},
	{0x55, 4},
	{0x7f, 7},
	{0xff, 8},
	{0x100, 1},
	{0x101, 2},
	{0xffff, 16},
	{0x10000, 1},
	{0x10001, 2},
	{0xffffffff, 32},
	{0x1ffffffff, 33},
	{0x3ffffffff, 34},
	{0x7ffffffff, 35},
	{0xfffffffff, 36},
	{0x3fffffffffffffff, 62},
	{0x7fffffffffffffff, 63},
	{0xffffffffffffffff, 64},
}

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

var y int = 2

func BenchmarkCountBitsInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsInt8(int8(i))
	}
}

func BenchmarkCountBitsInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsInt16(int16(i))
	}
}

func BenchmarkCountBitsInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsInt32(int32(i))
	}
}

func BenchmarkCountBitsInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsInt64(int64(i))
	}
}

func BenchmarkCountBitsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsInt(int(i))
	}
}

func BenchmarkCountBitsUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsUint16(uint16(i))
	}
}

func BenchmarkCountBitsUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsUint32(uint32(i))
	}
}

func BenchmarkCountBitsUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsUint64(uint64(i))
	}
}

func BenchmarkCountBitsUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsUint(uint(i))
	}
}

func BenchmarkCountBitsByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsByte(byte(i))
	}
}

func BenchmarkCountBitsByteAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsByteAlt(byte(i))
	}
}

func BenchmarkCountBitsRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y += CountBitsRune(rune(i))
	}
}

func BenchmarkDone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second)
	}
	b.Logf("print to prevent dead-code elimination %d", y)
}
