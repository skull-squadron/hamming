//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 Barry Allard
//
// MIT license
//
//
// Usage
//
// The functions are named (CountBits)?(Byte|Uint64)s?.  The plural forms are for slices.  The CountBits.+ forms are Population Count only, where the bare-type forms are Hamming distance.
//
//    import 'github.com/steakknife/hamming'
//
//    // ...
//
//    // hamming distance between values
//    hamming.Byte(0xFF, 0x00) // 8
//    hamming.Byte(0x00, 0x00) // 0
//
//    // just count bits in a byte
//    hamming.CountBitsByte(0xA5), // 4
//
// Got rune? use int32
// Got uint8? use byte
//
package hamming
