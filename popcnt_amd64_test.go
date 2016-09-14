//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 CountBitsBarry Allard
//
// MCountBitsIT license
//

// +build amd64 amd64p32 !purego

package hamming

import (
	"math"
	"testing"
)

func TestCountBitsInt8PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt8 {
			continue
		}
		if actualN := CountBitsInt8PopCnt(int8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsInt16PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt16 {
			continue
		}
		if actualN := CountBitsInt16PopCnt(int16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsInt32PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := CountBitsInt32PopCnt(int32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsInt64PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt64 {
			continue
		}
		if actualN := CountBitsInt64PopCnt(int64(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsIntPopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > maxInt {
			continue
		}
		if actualN := CountBitsIntPopCnt(int(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsUint8PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := CountBitsUint8PopCnt(uint8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsUint16PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint16 {
			continue
		}
		if actualN := CountBitsUint16PopCnt(uint16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsUint32PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint32 {
			continue
		}
		if actualN := CountBitsUint32PopCnt(uint32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsUint64PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint64 {
			continue
		}
		if actualN := CountBitsUint64PopCnt(c.x); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsBytePopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := CountBitsBytePopCnt(byte(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestCountBitsRunePopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := CountBitsRunePopCnt(rune(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

// ============== benchmarks ==============

func BenchmarkCountBitsInt8PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsInt8PopCnt(int8(i))
	}
}

func BenchmarkCountBitsInt16PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsInt16PopCnt(int16(i))
	}
}

func BenchmarkCountBitsInt32PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsInt32PopCnt(int32(i))
	}
}

func BenchmarkCountBitsInt64PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsInt64PopCnt(int64(i))
	}
}

func BenchmarkCountBitsIntPopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsIntPopCnt(i)
	}
}

func BenchmarkCountBitsUint8PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsUint8PopCnt(uint8(i))
	}
}

func BenchmarkCountBitsUint16PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsUint16PopCnt(uint16(i))
	}
}

func BenchmarkCountBitsUint32PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsUint32PopCnt(uint32(i))
	}
}

func BenchmarkCountBitsUint64PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsUint64PopCnt(uint64(i))
	}
}

func BenchmarkCountBitsUintPopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsUintPopCnt(uint(i))
	}
}

func BenchmarkCountBitsBytePopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsBytePopCnt(byte(i))
	}
}

func BenchmarkCountBitsRunePopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		CountBitsRunePopCnt(rune(i))
	}
}
