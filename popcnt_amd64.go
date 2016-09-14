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

func HasPopCnt() (ret bool)

func PopCntInt8(x int8) (ret int)
func PopCntInt16(x int16) (ret int)
func PopCntInt32(x int32) (ret int)
func PopCntInt64(x int64) (ret int)
func PopCntInt(x int) int {
	if strconv.IntSize == 64 {
		return PopCntInt64(int64(x))
	} else if strconv.IntSize == 32 {
		return PopCntInt32(int32(x))
	}
	panic("strconv.IntSize must be 32 or 64")
}
func PopCntUint8(x uint8) (ret int)
func PopCntUint16(x uint16) (ret int)
func PopCntUint32(x uint32) (ret int)
func PopCntUint64(x uint64) (ret int)
func PopCntUint(x uint) int {
	if strconv.IntSize == 64 {
		return PopCntUint64(uint64(x))
	} else if strconv.IntSize == 32 {
		return PopCntUint32(uint32(x))
	}
	panic("strconv.IntSize must be 32 or 64")
}
func PopCntByte(x byte) (ret int)
func PopCntRune(x rune) (ret int)
