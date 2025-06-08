// Copyright 2020 Justine Alexandra Roberts Tunney
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#define ULAW_BIAS $0x84

// func L16MixSat160(dst, src *int16)
// Explanation: http://i.imgur.com/nejgQ41.jpg
TEXT    ·L16MixSat160(SB),4,$0-16
	MOVQ	dst+0(FP), AX
	MOVQ	src+8(FP), BX
	MOVQ	$19, CX
moar:	MOVO	0(BX), X0
	PADDSW  0(AX), X0
	MOVO	X0, 0(AX)
	ADDQ	$16, AX
	ADDQ	$16, BX
	DECQ	CX
	CMPQ	CX, $0
	JGE	moar
	RET

// func LinearToUlaw(linear int64) (ulaw int64)
TEXT	·LinearToUlaw(SB),4,$0-16
	MOVQ	linear+0(FP), AX
	CMPQ	AX, $0		// if rax < 0: (Jump if Signed)
	JS	ifNeg		//   goto ifNeg
ifPos:	ADDQ	ULAW_BIAS, AX
	MOVQ	$0xFF, R13	// mask bits we'll apply later
	JMP	endIf
ifNeg:	MOVQ	ULAW_BIAS, BX	// rax = ULAW_BIAS - rax
	SUBQ	AX, BX
	MOVQ	BX, AX
	MOVQ	$0x7F, R13	// mask bits we'll apply later
endIf:	MOVQ	AX, DX
	ORQ	$0xFF, DX
	BSRQ	DX, BX		// get position of highest bit
	SUBQ	$7, BX
	CMPQ	BX, $8		// if rcx >= 8:
	JGE	loud		//   goto loud
	MOVQ	BX, CX
	ADDQ	$3, CX
	SARQ	CX, AX
	ANDQ	$0xF, AX
	SHLQ	$4, BX
	ORQ	BX, AX
	XORQ	R13, AX
	MOVQ	AX, ulaw+8(FP)
	RET
loud:	XORQ	$0x7F, R13
	MOVQ	R13, ulaw+8(FP)
	RET

// func UlawToLinear(ulaw int64) (linear int64)
TEXT	·UlawToLinear(SB),4,$0-16
	MOVQ	ulaw+0(FP), AX
	NOTQ	AX
	MOVQ	AX, CX
	MOVQ	AX, DX
	ANDQ	$0x0F, AX
	SHLQ	$3, AX
	ADDQ	$0x84, AX
	ANDQ	$0x70, CX
	SARQ	$4, CX
	SHLQ	CX, AX
	ANDQ	$0x80, DX
	CMPQ	DX, $0
	JE	death
	MOVQ	$0x84, BX
	SUBQ	AX, BX
	MOVQ	BX, linear+8(FP)
	RET
death:	SUBQ	$0x84, AX
	MOVQ	AX, linear+8(FP)
	RET
