package hamming

// TODO: SSE4.x PopCnt

const (
  m1  uint64 = 0x5555555555555555 //binary: 0101...
  m2  uint64 = 0x3333333333333333 //binary: 00110011..
  m4  uint64 = 0x0f0f0f0f0f0f0f0f //binary:  4 zeros,  4 ones ...
  m8  uint64 = 0x00ff00ff00ff00ff //binary:  8 zeros,  8 ones ...
  m16 uint64 = 0x0000ffff0000ffff //binary: 16 zeros, 16 ones ...
  m32 uint64 = 0x00000000ffffffff //binary: 32 zeros, 32 ones
  hff uint64 = 0xffffffffffffffff //binary: all ones
  h01 uint64 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3...
)

// hamming distance of two uint64's
func Uint64(x, y uint64) int {
  return CountBitsUint64(x ^ y)
}

// hamming distance of two bytes
func Byte(x, y byte) int {
  return CountBitsByte(x ^ y)
}

func CountBitsUint64(x uint64) int {
  x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
  x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
  x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
  return int((x * h01) >> 56)    //returns left 8 bits of x + (x<<8) + (x<<16) + (x<<24) + ...
}

func CountBitsByte(x byte) int {
  c := x - ((x >> 1) & 0333) - ((x >> 2) & 0111)
  return int((c+(c>>3))&0307) % 63
}
