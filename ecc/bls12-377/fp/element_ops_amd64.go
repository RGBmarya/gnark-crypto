//go:build !purego
// +build !purego

// Copyright 2020 ConsenSys Software Inc.
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

// Code generated by consensys/gnark-crypto DO NOT EDIT

package fp

import (
	"math/bits"
	"golang.org/x/exp/slices"
)

//go:noescape
func MulBy3(x *Element)

//go:noescape
func MulBy5(x *Element)

//go:noescape
func MulBy13(x *Element)

//go:noescape
func mul(res, x, y *Element)

//go:noescape
func fromMont(res *Element)

//go:noescape
func reduce(res *Element)

// Butterfly sets
//
//	a = a + b (mod q)
//	b = a - b (mod q)
//
//go:noescape
func Butterfly(a, b *Element)

// Mul z = x * y (mod q)
//
// x and y must be less than q
func (z *Element) MulCIOS(x, y *Element) *Element {

	// Implements CIOS multiplication -- section 2.3.2 of Tolga Acar's thesis
	// https://www.microsoft.com/en-us/research/wp-content/uploads/1998/06/97Acar.pdf
	//
	// The algorithm:
	//
	// for i=0 to N-1
	// 		C := 0
	// 		for j=0 to N-1
	// 			(C,t[j]) := t[j] + x[j]*y[i] + C
	// 		(t[N+1],t[N]) := t[N] + C
	//
	// 		C := 0
	// 		m := t[0]*q'[0] mod D
	// 		(C,_) := t[0] + m*q[0]
	// 		for j=1 to N-1
	// 			(C,t[j-1]) := t[j] + m*q[j] + C
	//
	// 		(C,t[N-1]) := t[N] + C
	// 		t[N] := t[N+1] + C
	//
	// → N is the number of machine words needed to store the modulus q
	// → D is the word size. For example, on a 64-bit architecture D is 2	64
	// → x[i], y[i], q[i] is the ith word of the numbers x,y,q
	// → q'[0] is the lowest word of the number -q⁻¹ mod r. This quantity is pre-computed, as it does not depend on the inputs.
	// → t is a temporary array of size N+2
	// → C, S are machine words. A pair (C,S) refers to (hi-bits, lo-bits) of a two-word number
	//
	// As described here https://hackmd.io/@gnark/modular_multiplication we can get rid of one carry chain and simplify:
	// (also described in https://eprint.iacr.org/2022/1400.pdf annex)
	//
	// for i=0 to N-1
	// 		(A,t[0]) := t[0] + x[0]*y[i]
	// 		m := t[0]*q'[0] mod W
	// 		C,_ := t[0] + m*q[0]
	// 		for j=1 to N-1
	// 			(A,t[j])  := t[j] + x[j]*y[i] + A
	// 			(C,t[j-1]) := t[j] + m*q[j] + C
	//
	// 		t[N-1] = C + A
	//
	// This optimization saves 5N + 2 additions in the algorithm, and can be used whenever the highest bit
	// of the modulus is zero (and not all of the remaining bits are set).

	mul(z, x, y)
	return z
}

//go:noescape
func VecMul_AVX2_I64(x []uint64, y []uint64, z []uint64, u []uint64)

//go:noescape
func VecAdd_AVX2_I64(x []uint64, y []uint64, z []uint64, u []uint64)

// Mihir
func VecAdd(x, y, carry []uint64) (sum0, sum1, carry0, carry1 uint64) {
	sum := make([]uint64, len(x))
	carryOut := slices.Clone(carry)
	VecAdd_AVX2_I64(x, y, carryOut, sum)
	return sum[0], sum[1], carryOut[0], carryOut[1]
}

func VecMul(x, y []uint64) (hi0, hi1, lo0, lo1 uint64) {
	hi := make([]uint64, len(x))
	low := make([]uint64, len(x))
	VecMul_AVX2_I64(x, y, hi, low)
	return hi[0], hi[1], low[0], low[1]
}

// Mul z = x * y (mod q)
//
// x and y must be less than q
func (c *Element) Mul(x, y *Element) *Element {
	// Implements a parallel radix-2^64 interleaved Montgomery multiplication algorithm
	// described in Algorithm 4 of the paper "Improved Montgomery Multiplication on SIMD Architectures"
	// https://eprint.iacr.org/2017/1057.pdf
	//
	// This algorithm is suitable for 32-bit 2-way SIMD vector instruction units, but our implementation
	// adapts it for 64-bit 2-way SIMD vector instruction units.
	//
	// The algorithm involves two main computations performed in parallel:
	//
	// Computation 1:
	// for j = 0 to n-1
	//     for i = 0 to n-1
	//         d_i = 0
	//         t_0 = a_j * b_0 + d_0
	//         t_0 = t_0 / 2^64
	//         for i = 1 to n-1
	//             p_0 = a_j * b_i + t_0 + d_i
	//             t_0 = p_0 / 2^64
	//             d_i-1 = p_0 mod 2^64
	//         end for
	//         d_n-1 = t_0
	// end for
	//
	// Computation 2:
	// for j = 0 to n-1
	//     for i = 0 to n-1
	//         e_i = 0
	//         q = ((μb_0)a_j + μ(d_0 - e_0)) mod 2^64
	//         t_1 = q * m_0 + e_0
	//         t_1 = t_1 / 2^64
	//         for i = 1 to n-1
	//             p_1 = q * m_i + t_1 + e_i
	//             t_1 = p_1 / 2^64
	//             e_i-1 = p_1 mod 2^64
	//         end for
	//         e_n-1 = t_1
	// end for
	//
	// The final result is calculated as:
	// C = D - E
	// where D = Σ(d_i * 2^(64i)) and E = Σ(e_i * 2^(64i))
	//
	// If C < 0 then C = C + M
	//
	// The notation used:
	// - a, b are the input operands
	// 		- parameters x, y represent a, b respectively
	// - m is the modulus
	// 		- q0...q5 in Mul represent the 'digits' of field modulus q
	// - n is the number of 64-bit words in the operands and modulus
	// 		- n = 6 in the case of fp
	// - μ is the precomputed constant -m^(-1) mod 2^64
	// 		- 'qInv' represents μ in Mul
	// - t_0, t_1 are temporary variables used for intermediate calculations
	// - p_0, p_1 are partial products
	// - d_i, e_i are intermediate results stored in arrays D and E
	//
	// This approach leverages SIMD vector instruction units to parallelize operations,
	// improving the efficiency of Montgomery multiplication on modern 64-bit architectures.

	// di = 0, ei = 0 for 0 <= i < n, where n is 6 in the case of fp
	var d0, d1, d2, d3, d4, d5 uint64
	var e0, e1, e2, e3, e4, e5 uint64
	

	{
		// first iteration -> j=0
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[0] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}

	{
		// second iteration -> j =1
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[1] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}

	{
		// third iteration -> j = 2
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[2] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}

	{
		// fourth iteration -> j = 3
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[3] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}
	
	{
		// fifth iteration -> j = 4
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[4] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}

	{
		// sixth iteration -> j = 5
		var t0, t1 uint64
		var diff_d0_e0, lo_aj_b0, sum_lo_ajb0_diff_d0e0 uint64 // temp vars for calculating q
		var lo_qm0 uint64 // temp vars for i = 0
		var hi_p0, hi_p1, lo_p0, lo_p1, sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei uint64 // temp vars for i = 1 ... (n - 1)
		var c1, c2 uint64

		aj := x[5] // x[j] for the j-th iteration
		b0 := y[0]

		// Calculating q
		// This is q is NOT the field modulus; q stores the intermediate value from Computation 2
		diff_d0_e0, _ = bits.Sub64(d0, e0, 0)
		// To avoid repeated computation of ajb0, we directly assign to t0 (Computation 1)
		// We operate on the lower 64 bits of q; mod(2^64) means we can ignore the upper 64 bits
		t0, lo_aj_b0 = bits.Mul64(aj, b0) 
		sum_lo_ajb0_diff_d0e0, _ = bits.Add64(lo_aj_b0, diff_d0_e0, 0)
		_, q := bits.Mul64(qInv, sum_lo_ajb0_diff_d0e0) 

		// i = 0 - this precedes the for loop
		t1, lo_qm0 = bits.Mul64(q, q0) //m_i in Algorithm 4 is qi here
		_, _, c1, c2 = VecAdd([]uint64{lo_aj_b0, lo_qm0}, []uint64{d0, e0}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{t0, t1}, []uint64{c1, c2}, []uint64{0, 0})

		// i = 1
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[1], q1})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
		d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 2
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[2], q2})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d2, e2}, []uint64{0, 0})
		d1, e1, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 3
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[3], q3})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d3, e3}, []uint64{0, 0})
		d2, e2, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 4
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[4], q4})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d4, e4}, []uint64{0, 0})
		d3, e3, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// i = 5
		hi_p0, hi_p1, lo_p0, lo_p1 = VecMul([]uint64{aj, q}, []uint64{y[5], q5})
		sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d5, e5}, []uint64{0, 0})
		d4, e4, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
		t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})

		// Final assignment after for loop with index i ends
		d5 = t0
		e5 = t1
	}

	var b uint64
	c[0], b = bits.Sub64(d0, e0, 0)
	c[1], b = bits.Sub64(d1, e1, b)
	c[2], b = bits.Sub64(d2, e2, b)
	c[3], b = bits.Sub64(d3, e3, b)
	c[4], b = bits.Sub64(d4, e4, b)
	c[5], b = bits.Sub64(d5, e5, b)

	if b == 1 {
		var carry uint64
		c[0], carry = bits.Add64(q0, c[0], 0)
		c[1], carry = bits.Add64(q1, c[1], carry)
		c[2], carry = bits.Add64(q2, c[2], carry)
		c[3], carry = bits.Add64(q3, c[3], carry)
		c[4], carry = bits.Add64(q4, c[4], carry)
		c[5], carry = bits.Add64(q5, c[5], carry)
	}
	return c
}

// Square z = x * x (mod q)
//
// x must be less than q
func (z *Element) Square(x *Element) *Element {
	// see Mul for doc.
	mul(z, x, x)
	return z
}
