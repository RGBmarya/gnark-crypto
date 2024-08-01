// Code generated by command: go run gen.go -out ../internal/functions/accel_avx2_amd64.s -stubs ../internal/functions/accel_avx2_amd64.go -pkg functions. DO NOT EDIT.

#include "textflag.h"

// func VecMul_AVX2_I64(x []uint64, y []uint64, z []uint64, u []uint64)
// Requires: AVX, AVX2
TEXT ·VecMul_AVX2_I64(SB), NOSPLIT, $0-96
	MOVQ     x_base+0(FP), DI
	MOVQ     y_base+24(FP), SI
	MOVQ     z_base+48(FP), DX
	MOVQ     u_base+72(FP), CX
	VMOVDQU  (DI), X0
	VMOVDQU  (SI), X1
	VPSRLQ   $0x20, X0, X2
	VPMULUDQ X0, X1, X3
	VPSRLQ   $0x20, X3, X4
	VPMULUDQ X2, X1, X5
	VPADDQ   X5, X4, X4
	VPXOR    X6, X6, X6
	VPBLENDD $0x0a, X6, X4, X6
	VPSRLQ   $0x20, X4, X4
	VPSRLQ   $0x20, X1, X1
	VPMULUDQ X0, X1, X0
	VPADDQ   X0, X6, X6
	VPMULUDQ X2, X1, X1
	VPADDQ   X1, X4, X1
	VPSRLQ   $0x20, X6, X2
	VPADDQ   X2, X1, X1
	VMOVDQU  X1, (DX)
	VPADDQ   X0, X5, X0
	VPSLLQ   $0x20, X0, X0
	VPADDQ   X0, X3, X0
	VMOVDQU  X0, (CX)
	RET

// func VecAdd_AVX2_I64(x []uint64, y []uint64, z []uint64, u []uint64)
// Requires: AVX
TEXT ·VecAdd_AVX2_I64(SB), NOSPLIT, $0-96
	MOVQ    x_base+0(FP), DI
	MOVQ    y_base+24(FP), SI
	MOVQ    z_base+48(FP), DX
	MOVQ    u_base+72(FP), CX
	VMOVDQA (DI), X0
	VMOVDQA (SI), X1
	VPADDQ  X0, X1, X2
	VPADDQ  (DX), X2, X2
	VMOVDQU X2, (CX)
	VPAND   X0, X1, X3
	VPOR    X0, X1, X0
	VPANDN  X0, X2, X0
	VPOR    X3, X0, X0
	VPSRLQ  $0x3f, X0, X0
	VMOVDQU X0, (DX)
	RET
