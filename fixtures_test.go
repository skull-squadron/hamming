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
	{0xfe, 7},
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
	{0x7ffffffffffffffe, 62},
	{0x7fffffffffffffff, 63},
	{0x8000000000000000, 1},
	{0x8000000000000001, 2},
	{0xffffffffffffffff, 64},
}
