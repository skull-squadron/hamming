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

TEXT ·PopCntInt8(SB),NOSPLIT,$0
	JMP        ·PopCntByte(SB)

TEXT ·PopCntInt16(SB),NOSPLIT,$0
	JMP        ·PopCntUint16(SB)

TEXT ·PopCntInt32(SB),NOSPLIT,$0
	JMP        ·PopCntUint32(SB)

TEXT ·PopCntInt64(SB),NOSPLIT,$0
	JMP        ·PopCntUint64(SB)

TEXT ·PopCntByte(SB),NOSPLIT,$0
	JMP        ·PopCntUint8(SB)

TEXT ·PopCntRune(SB),NOSPLIT,$0
	JMP        ·PopCntUint32(SB)

TEXT ·PopCntUint8(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVB       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·PopCntUint16(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVW       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·PopCntUint32(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVL       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·PopCntUint64(SB),NOSPLIT,$0
	POPCNTQ    x+0(FP), AX
	MOVQ       AX, ret+8(FP)
	RET

// func hasPopCnt() (ret bool) 
TEXT ·HasPopCnt(SB),NOSPLIT,$0
	MOVL       $1, AX
	CPUID
	SHRL       $23, CX  // bit 23: Advanced Bit Manipulation Bit (ABM) -> POPCNTQ
	ANDL       $1, CX
	MOVB       CX, ret+0(FP)
	RET
