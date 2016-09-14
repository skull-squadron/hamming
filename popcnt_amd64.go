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

func Int8PopCnt(x int8) (ret int)
func Int16PopCnt(x int16) (ret int)
func Int32PopCnt(x int32) (ret int)
func Int64PopCnt(x int64) (ret int)
func IntPopCnt(x int) int {
	if strconv.IntSize == 64 {
		return Int64PopCnt(int64(x))
	} else if strconv.IntSize == 32 {
		return Int32PopCnt(int32(x))
	}
	panic("strconv.IntSize must be 32 or 64")
}
func Uint8PopCnt(x uint8) (ret int)
func Uint16PopCnt(x uint16) (ret int)
func Uint32PopCnt(x uint32) (ret int)
func Uint64PopCnt(x uint64) (ret int)
func UintPopCnt(x uint) int {
	if strconv.IntSize == 64 {
		return Uint64PopCnt(uint64(x))
	} else if strconv.IntSize == 32 {
		return Uint32PopCnt(uint32(x))
	}
	panic("strconv.IntSize must be 32 or 64")
}
func BytePopCnt(x byte) (ret int)
func RunePopCnt(x rune) (ret int)
