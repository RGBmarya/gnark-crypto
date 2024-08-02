// Code generated by command: go run gen.go -out ../internal/functions/accel_avx2_amd64.s -stubs ../internal/functions/accel_avx2_amd64.go -pkg functions. DO NOT EDIT.

#include "textflag.h"

// func VecMul_I64_AVX2(x []int, y []int, z []int, u []int)
// Requires: AVX, AVX2
TEXT ·VecMul_I64_AVX2(SB), NOSPLIT, $8-96
	MOVQ  x_base+0(FP), DI
	MOVQ  y_base+24(FP), SI
	MOVQ  z_base+48(FP), DX
	MOVQ  u_base+72(FP), CX
	PUSHQ BP
	MOVQ  SP, BP
	SUBQ  $0x20, SP
	MOVQ  DI, -8(BP)
	MOVQ  SI, -16(BP)
	MOVQ  DX, -24(BP)
	MOVQ  CX, -32(BP)
	MOVQ  -32(BP), CX
	MOVQ  -24(BP), DX
	MOVQ  -16(BP), SI
	MOVQ  -8(BP), AX
	MOVQ  AX, DI
	CALL  VecMul_v
	NOP
	RET

// func VecAdd_I64_AVX2(x []int, y []int, z []int, u []int)
// Requires: AVX, SSE, SSE2
TEXT ·VecAdd_I64(SB), NOSPLIT, $8-96
	MOVQ  x_base+0(FP), DI
	MOVQ  y_base+24(FP), SI
	MOVQ  z_base+48(FP), DX
	MOVQ  u_base+72(FP), CX
	PUSHQ BP
	MOVQ  SP, BP
	SUBQ  $0x20, SP
	MOVQ  DI, -8(BP)
	MOVQ  SI, -16(BP)
	MOVQ  DX, -24(BP)
	MOVQ  CX, -32(BP)
	MOVQ  -32(BP), CX
	MOVQ  -24(BP), DX
	MOVQ  -16(BP), SI
	MOVQ  -8(BP), AX
	MOVQ  AX, DI
	CALL  VecAdd_v
	NOP
	RET

// func VecMul_V(x []int, y []int, z []int, u []int)
TEXT ·VecMul_V(SB), NOSPLIT, $8-96
	MOVQ  x_base+0(FP), DI
	MOVQ  y_base+24(FP), SI
	MOVQ  z_base+48(FP), DX
	MOVQ  u_base+72(FP), CX
	PUSHQ BP
	MOVQ  SP, BP
	MOVQ  DI, -88(BP)
	MOVQ  SI, -96(BP)
	MOVQ  DX, -104(BP)
	MOVQ  CX, -112(BP)
	MOVL  $0xffffffff, AX
	MOVQ  AX, -8(BP)
	MOVQ  -88(BP), AX
	MOVQ  (AX), AX
	ANDL  $0xffffffff, AX
	MOVQ  AX, -16(BP)
	MOVQ  -88(BP), AX
	MOVQ  (AX), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -24(BP)
	MOVQ  -96(BP), AX
	MOVQ  (AX), AX
	ANDL  $0xffffffff, AX
	MOVQ  AX, -32(BP)
	MOVQ  -96(BP), AX
	MOVQ  (AX), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -40(BP)
	MOVQ  -16(BP), AX
	IMULQ -32(BP), AX
	MOVQ  AX, -48(BP)
	MOVQ  -24(BP), AX
	IMULQ -32(BP), AX
	MOVQ  -48(BP), DX
	SHRQ  $0x20, DX
	ADDQ  DX, AX
	MOVQ  AX, -56(BP)
	MOVQ  -56(BP), AX
	MOVQ  AX, -64(BP)
	MOVQ  -56(BP), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -72(BP)
	MOVQ  -16(BP), AX
	IMULQ -40(BP), AX
	ADDQ  AX, -64(BP)
	MOVQ  -24(BP), AX
	IMULQ -40(BP), AX
	MOVQ  AX, DX
	MOVQ  -72(BP), AX
	ADDQ  AX, DX
	MOVQ  -64(BP), AX
	SHRQ  $0x20, AX
	ADDQ  AX, DX
	MOVQ  -104(BP), AX
	MOVQ  DX, (AX)
	MOVQ  -88(BP), AX
	MOVQ  (AX), DX
	MOVQ  -96(BP), AX
	MOVQ  (AX), AX
	IMULQ AX, DX
	MOVQ  -112(BP), AX
	MOVQ  DX, (AX)
	MOVQ  -88(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	ANDL  $0xffffffff, AX
	MOVQ  AX, -16(BP)
	MOVQ  -88(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -24(BP)
	MOVQ  -96(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	ANDL  $0xffffffff, AX
	MOVQ  AX, -32(BP)
	MOVQ  -96(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -40(BP)
	MOVQ  -16(BP), AX
	IMULQ -32(BP), AX
	MOVQ  AX, -48(BP)
	MOVQ  -24(BP), AX
	IMULQ -32(BP), AX
	MOVQ  -48(BP), DX
	SHRQ  $0x20, DX
	ADDQ  DX, AX
	MOVQ  AX, -56(BP)
	MOVQ  -56(BP), AX
	MOVQ  AX, -64(BP)
	MOVQ  -56(BP), AX
	SHRQ  $0x20, AX
	MOVQ  AX, -72(BP)
	MOVQ  -16(BP), AX
	IMULQ -40(BP), AX
	ADDQ  AX, -64(BP)
	MOVQ  -24(BP), AX
	IMULQ -40(BP), AX
	MOVQ  AX, DX
	MOVQ  -72(BP), AX
	LEAQ  (DX)(AX*1), CX
	MOVQ  -64(BP), AX
	SHRQ  $0x20, AX
	MOVQ  AX, DX
	MOVQ  -104(BP), AX
	ADDQ  $0x08, AX
	ADDQ  CX, DX
	MOVQ  DX, (AX)
	MOVQ  -88(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), DX
	MOVQ  -96(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	MOVQ  -112(BP), CX
	ADDQ  $0x08, CX
	IMULQ DX, AX
	MOVQ  AX, (CX)
	NOP
	POPQ  BP
	RET

// func VecAdd_V(x []int, y []int, z []int, u []int)
TEXT ·VecAdd_V(SB), NOSPLIT, $8-96
	MOVQ  x_base+0(FP), DI
	MOVQ  y_base+24(FP), SI
	MOVQ  z_base+48(FP), DX
	MOVQ  u_base+72(FP), CX
	PUSHQ BP
	MOVQ  SP, BP
	MOVQ  DI, -8(BP)
	MOVQ  SI, -16(BP)
	MOVQ  DX, -24(BP)
	MOVQ  CX, -32(BP)
	MOVQ  -8(BP), AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	MOVQ  (AX), AX
	ADDQ  AX, DX
	MOVQ  -24(BP), AX
	MOVQ  (AX), AX
	ADDQ  AX, DX
	MOVQ  -32(BP), AX
	MOVQ  DX, (AX)
	MOVQ  -8(BP), AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	MOVQ  (AX), AX
	MOVQ  DX, CX
	ANDQ  AX, CX
	MOVQ  -8(BP), AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	MOVQ  (AX), AX
	ORQ   AX, DX
	MOVQ  -32(BP), AX
	MOVQ  (AX), AX
	NOTQ  AX
	ANDQ  DX, AX
	ORQ   CX, AX
	SHRQ  $0x3f, AX
	MOVQ  AX, DX
	MOVQ  -24(BP), AX
	MOVQ  DX, (AX)
	MOVQ  -8(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	LEAQ  (DX)(AX*1), CX
	MOVQ  -24(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), DX
	MOVQ  -32(BP), AX
	ADDQ  $0x08, AX
	ADDQ  CX, DX
	MOVQ  DX, (AX)
	MOVQ  -8(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	MOVQ  DX, CX
	ANDQ  AX, CX
	MOVQ  -8(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), DX
	MOVQ  -16(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	ORQ   AX, DX
	MOVQ  -32(BP), AX
	ADDQ  $0x08, AX
	MOVQ  (AX), AX
	NOTQ  AX
	ANDQ  DX, AX
	ORQ   AX, CX
	MOVQ  CX, DX
	MOVQ  -24(BP), AX
	ADDQ  $0x08, AX
	SHRQ  $0x3f, DX
	MOVQ  DX, (AX)
	NOP
	POPQ  BP
	RET
