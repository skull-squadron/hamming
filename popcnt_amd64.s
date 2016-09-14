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

TEXT ·Int8PopCnt(SB),NOSPLIT,$0
	JMP        ·BytePopCnt(SB)

TEXT ·Int16PopCnt(SB),NOSPLIT,$0
	JMP        ·Uint16PopCnt(SB)

TEXT ·Int32PopCnt(SB),NOSPLIT,$0
	JMP        ·Uint32PopCnt(SB)

TEXT ·Int64PopCnt(SB),NOSPLIT,$0
	JMP        ·Uint64PopCnt(SB)

TEXT ·BytePopCnt(SB),NOSPLIT,$0
	JMP        ·Uint8PopCnt(SB)

TEXT ·RunePopCnt(SB),NOSPLIT,$0
	JMP        ·Uint32PopCnt(SB)

TEXT ·Uint8PopCnt(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVB       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·Uint16PopCnt(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVW       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·Uint32PopCnt(SB),NOSPLIT,$0
	XORQ       AX, AX
	MOVL       x+0(FP), AX
	POPCNTQ    AX, AX	
	MOVQ       AX, ret+8(FP)
	RET

TEXT ·Uint64PopCnt(SB),NOSPLIT,$0
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
