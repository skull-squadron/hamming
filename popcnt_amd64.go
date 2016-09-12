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

func PopCntInt8(x int8) (ret int32)
func PopCntInt16(x int16) (ret int32)
func PopCntInt32(x int32) (ret int32)
func PopCntInt64(x int64) (ret int32)

func PopCntUint8(x uint8) (ret int32)
func PopCntUint16(x uint16) (ret int32)
func PopCntUint32(x uint32) (ret int32)
func PopCntUint64(x uint64) (ret int32)

func PopCntByte(x byte) (ret int32)
func PopCntRune(x rune) (ret int32)

func PopCntUint(x uint) (ret int32) {
	switch strconv.IntSize {
	case 32:
		return PopCntUint32(uint32(x))
	case 64:
		return PopCntUint64(uint64(x))
	}
	panic("strconv.IntSize should be 32 or 64")
}

func PopCntInt(x int) (ret int32) {
	switch strconv.IntSize {
	case 32:
		return PopCntInt32(int32(x))
	case 64:
		return PopCntInt64(int64(x))
	}
	panic("strconv.IntSize should be 32 or 64")
}
