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

import "strconv"

// References: check out Hacker's Delight, about p. 70

var table = [256]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

// table-less, branch-free implementation
func CountBitsByteAlt(x byte) int {
	x = (x & 0x55) + ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	return int((x & 0x0f) + ((x >> 4) & 0x0f))
}

func CountBitsInt8(x int8) int   { return CountBitsByte(byte(x)) }
func CountBitsInt16(x int16) int { return CountBitsUint16(uint16(x)) }
func CountBitsInt32(x int32) int { return CountBitsUint32(uint32(x)) }
func CountBitsInt64(x int64) int { return CountBitsUint64(uint64(x)) }
func CountBitsInt(x int) int     { return CountBitsUint(uint(x)) }
func CountBitsByte(x byte) int   { return CountBitsUint8(x) }
func CountBitsRune(x rune) int   { return CountBitsInt32(x) }

func CountBitsUint8(x uint8) int {
	return int(table[x])
}

func CountBitsUint16(x uint16) int {
	return int(table[x&0xFF] + table[(x>>8)&0xFF])
}

const (
	m1d uint32 = 0x55555555
	m2d        = 0x33333333
	m4d        = 0x0f0f0f0f
)

func CountBitsUint32(x uint32) int {
	x -= ((x >> 1) & m1d)
	x = (x & m2d) + ((x >> 2) & m2d)
	x = (x + (x >> 4)) & m4d
	x += x >> 8
	x += x >> 16
	return int(x & 0x3f)
}

const (
	m1q uint64 = 0x5555555555555555
	m2q        = 0x3333333333333333
	m4q        = 0x0f0f0f0f0f0f0f0f
	hq         = 0x0101010101010101
)

func CountBitsUint64(x uint64) int {
	x -= (x >> 1) & m1q              // put count of each 2 bits into those 2 bits
	x = (x & m2q) + ((x >> 2) & m2q) // put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4q         // put count of each 8 bits into those 8 bits
	return int((x * hq) >> 56)       // returns left 8 bits of x + (x<<8) + (x<<16) + (x<<24) + ...
}

func CountBitsUint64Alt(x uint64) int {
	return CountBitsUint32(uint32(x>>32)) + CountBitsUint32(uint32(x))
}

func CountBitsUintReference(x uint) int {
	c := 0
	for x != 0 {
		x &= x - 1
		c++
	}
	return c
}

func CountBitsUint(x uint) int {
	return CountBitsUint64(uint64(x))
}

var assertIntSizeNotMoreThan64Bits = func() interface{} {
	if strconv.IntSize > 64 {
		panic("CountBitsUint cannot function when IntSize > 64 bits")
	}
	return nil
}()
