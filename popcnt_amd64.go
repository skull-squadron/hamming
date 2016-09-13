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

func hasPopCnt() (ret bool)

var (
	PopCntInt8   func(x int8) int
	PopCntInt16  func(x int16) int
	PopCntInt32  func(x int32) int
	PopCntInt64  func(x int64) int
	PopCntUint8  func(x uint8) int
	PopCntUint16 func(x uint16) int
	PopCntUint32 func(x uint32) int
	PopCntUint64 func(x uint64) int
	PopCntByte   func(x byte) int
	PopCntRune   func(x rune) int
)

var setup = func() interface{} {
	if hasPopCnt() {
		PopCntInt8 = popCntInt8
		PopCntInt16 = popCntInt16
		PopCntInt32 = popCntInt32
		PopCntInt64 = popCntInt64
		PopCntUint8 = popCntUint8
		PopCntUint16 = popCntUint16
		PopCntUint32 = popCntUint32
		PopCntUint64 = popCntUint64
		PopCntByte = popCntByte
		PopCntRune = popCntRune
	} else { // no POPCNTQ, fallback to regular functions
		PopCntInt8 = CountBitsInt8
		PopCntInt16 = CountBitsInt16
		PopCntInt32 = CountBitsInt32
		PopCntInt64 = CountBitsInt64
		PopCntUint8 = CountBitsUint8
		PopCntUint16 = CountBitsUint16
		PopCntUint32 = CountBitsUint32
		PopCntUint64 = CountBitsUint64
		PopCntByte = CountBitsByte
		PopCntRune = CountBitsRune
	}
	return nil
}()

func PopCntUint(x uint) int {
	switch strconv.IntSize {
	case 32:
		return PopCntUint32(uint32(x))
	case 64:
		return PopCntUint64(uint64(x))
	}
	panic("strconv.IntSize should be 32 or 64")
}

func PopCntInt(x int) int {
	switch strconv.IntSize {
	case 32:
		return PopCntInt32(int32(x))
	case 64:
		return PopCntInt64(int64(x))
	}
	panic("strconv.IntSize should be 32 or 64")
}

func popCntInt8(x int8) (ret int)
func popCntInt16(x int16) (ret int)
func popCntInt32(x int32) (ret int)
func popCntInt64(x int64) (ret int)
func popCntUint8(x uint8) (ret int)
func popCntUint16(x uint16) (ret int)
func popCntUint32(x uint32) (ret int)
func popCntUint64(x uint64) (ret int)
func popCntByte(x byte) (ret int)
func popCntRune(x rune) (ret int)
