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

func TestInt8PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt8 {
			continue
		}
		if actualN := Int8PopCnt(int8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestInt16PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt16 {
			continue
		}
		if actualN := Int16PopCnt(int16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestInt32PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := Int32PopCnt(int32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestInt64PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt64 {
			continue
		}
		if actualN := Int64PopCnt(int64(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestIntPopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > maxInt {
			continue
		}
		if actualN := IntPopCnt(int(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestUint8PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := Uint8PopCnt(uint8(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestUint16PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint16 {
			continue
		}
		if actualN := Uint16PopCnt(uint16(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestUint32PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint32 {
			continue
		}
		if actualN := Uint32PopCnt(uint32(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestUint64PopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint64 {
			continue
		}
		if actualN := Uint64PopCnt(c.x); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestBytePopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxUint8 {
			continue
		}
		if actualN := BytePopCnt(byte(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

func TestRunePopCnt(t *testing.T) {
	if !HasPopCnt() {
		t.SkipNow()
	}
	for _, c := range testCountBitsCases {
		if c.x > math.MaxInt32 {
			continue
		}
		if actualN := RunePopCnt(rune(c.x)); actualN != c.n {
			t.Fatalf("%d -> (actual) %d != %d (expected)", c.x, actualN, c.n)
		}
	}
}

// ============== benchmarks ==============

func BenchmarkInt8PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Int8PopCnt(int8(i))
	}
}

func BenchmarkInt16PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Int16PopCnt(int16(i))
	}
}

func BenchmarkInt32PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Int32PopCnt(int32(i))
	}
}

func BenchmarkInt64PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Int64PopCnt(int64(i))
	}
}

func BenchmarkIntPopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		IntPopCnt(i)
	}
}

func BenchmarkUint8PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Uint8PopCnt(uint8(i))
	}
}

func BenchmarkUint16PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Uint16PopCnt(uint16(i))
	}
}

func BenchmarkUint32PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Uint32PopCnt(uint32(i))
	}
}

func BenchmarkUint64PopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		Uint64PopCnt(uint64(i))
	}
}

func BenchmarkUintPopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		UintPopCnt(uint(i))
	}
}

func BenchmarkBytePopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		BytePopCnt(byte(i))
	}
}

func BenchmarkRunePopCnt(b *testing.B) {
	if !HasPopCnt() {
		b.SkipNow()
	}
	for i := 0; i < b.N; i++ {
		RunePopCnt(rune(i))
	}
}
