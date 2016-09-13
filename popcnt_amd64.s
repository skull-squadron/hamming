//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 Barry Allard
//
// MIT license
//

#include "textflag.h"

TEXT ·popCntInt8(SB),NOSPLIT,$0
	JMP        ·popCntByte(SB)

TEXT ·popCntInt16(SB),NOSPLIT,$0
	JMP        ·popCntUint16(SB)

TEXT ·popCntInt32(SB),NOSPLIT,$0
	JMP        ·popCntUint32(SB)

TEXT ·popCntInt64(SB),NOSPLIT,$0
	JMP        ·popCntUint64(SB)

TEXT ·popCntByte(SB),NOSPLIT,$0
	JMP        ·popCntUint8(SB)

TEXT ·popCntRune(SB),NOSPLIT,$0
	JMP        ·popCntUint32(SB)

TEXT ·popCntUint8(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVB       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·popCntUint16(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVW       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·popCntUint32(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVL       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·popCntUint64(SB),NOSPLIT,$0
	POPCNTQ    x+0(FP), AX
	MOVQ       AX, ret+8(FP)
	RET

// func hasPopCnt() (ret bool) 
TEXT ·hasPopCnt(SB),NOSPLIT,$0
	MOVL       $1, AX
	CPUID
	SHRL       $23, CX  // bit 23: Advanced Bit Manipulation Bit (ABM) -> POPCNTQ
	ANDL       $1, CX
	MOVB       CX, ret+0(FP)
	RET
