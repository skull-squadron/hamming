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

import "testing"

type testBytesCase struct {
	b0, b1 []byte
	n      int
}

var testBytesCases = []testBytesCase{
	{[]byte{}, []byte{}, 0},
	{[]byte{1}, []byte{0}, 1},
	{[]byte{1}, []byte{2}, 2},
	{[]byte{1, 0}, []byte{0, 1}, 2},
	{[]byte{1, 0}, []byte{0, 1}, 2},
}

type testUint64sCase struct {
	b0, b1 []uint64
	n      int
}

var testUint64sCases = []testUint64sCase{
	{[]uint64{}, []uint64{}, 0},
	{[]uint64{1}, []uint64{0}, 1},
	{[]uint64{1}, []uint64{2}, 2},
	{[]uint64{1, 0}, []uint64{0, 1}, 2},
	{[]uint64{1, 0}, []uint64{0, 1}, 2},
}

func TestBytes(t *testing.T) {
	for _, c := range testBytesCases {
		if actualN := Bytes(c.b0, c.b1); actualN != c.n {
			t.Fatalf("Bytes(%d,%d) = %d != %d", c.b0, c.b1, actualN, c.n)
		} else {
			t.Logf("Bytes(%d,%d) = %d == %d", c.b0, c.b1, actualN, c.n)
		}
	}
}

func TestUint64s(t *testing.T) {
	for _, c := range testUint64sCases {
		if actualN := Uint64s(c.b0, c.b1); actualN != c.n {
			t.Fatalf("Uint64s(%d,%d) = %d != %d", c.b0, c.b1, actualN, c.n)
		} else {
			t.Logf("Uint64s(%d,%d) = %d == %d", c.b0, c.b1, actualN, c.n)
		}
	}
}
