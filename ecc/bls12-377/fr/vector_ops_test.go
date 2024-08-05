package fr

import (
	"fmt"
	"math/bits"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestVecAdd(t *testing.T) {
	// Test case 1: Vectors with zero elements
	x := []uint64{0, 0}
	y := []uint64{0, 0}
	carry := []uint64{0, 0}
	expectedSum0, expectedCarry0 := bits.Add64(x[0], y[0], carry[0])
	expectedSum1, expectedCarry1 := bits.Add64(x[1], y[1], carry[1])
	sum0, sum1, carry0, carry1 := VecAdd(x, y, carry)
	fmt.Println(sum0, sum1, carry0, carry1)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}

	// Test case 2: Vectors with single element
	x = []uint64{1, 0}
	y = []uint64{2, 0}
	carry = []uint64{0, 0}
	expectedSum0, expectedCarry0 = bits.Add64(x[0], y[0], carry[0])
	expectedSum1, expectedCarry1 = bits.Add64(x[1], y[1], carry[1])
	sum0, sum1, carry0, carry1 = VecAdd(x, y, carry)
	fmt.Println(sum0, sum1, carry0, carry1)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 2 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}

	// Test case 3: Vectors with multiple elements
	x = []uint64{1, 2}
	y = []uint64{4, 5}
	carry = []uint64{0, 0}
	expectedSum0, expectedCarry0 = bits.Add64(x[0], y[0], carry[0])
	expectedSum1, expectedCarry1 = bits.Add64(x[1], y[1], carry[1])
	sum0, sum1, carry0, carry1 = VecAdd(x, y, carry)
	fmt.Println(sum0, sum1, carry0, carry1)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 3 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}

	// Test case 4: Vectors with carry
	x = []uint64{1, 2}
	y = []uint64{4, 5}
	carry = []uint64{1, 1}
	expectedSum0, expectedCarry0 = bits.Add64(x[0], y[0], carry[0])
	expectedSum1, expectedCarry1 = bits.Add64(x[1], y[1], carry[1])
	sum0, sum1, carry0, carry1 = VecAdd(x, y, carry)
	fmt.Println(sum0, sum1, carry0, carry1)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 4 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}

	// Test case 5: Vectors with large numbers
	x = []uint64{18446744073709551615, 18446744073709551615}
	y = []uint64{1, 1}
	carry = []uint64{0, 0}
	expectedSum0, expectedCarry0 = bits.Add64(x[0], y[0], carry[0])
	expectedSum1, expectedCarry1 = bits.Add64(x[1], y[1], carry[1])
	sum0, sum1, carry0, carry1 = VecAdd(x, y, carry)
	fmt.Println(sum0, sum1, carry0, carry1)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 5 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}
}

func TestVecAddAligned(t *testing.T) {
	x := align32Uint64(2)
	y := align32Uint64(2)
	carry := align32Uint64(2)
	addr := uintptr(unsafe.Pointer(&x[0]))
	if addr%32 == 0 {
		fmt.Println("Array is 32-byte aligned")
	} else {
		fmt.Println("Array is not 32-byte aligned")
	}

	fmt.Printf("Aligned Array Address: %x\n", addr)
	addr = uintptr(unsafe.Pointer(&y[0]))
	if addr%32 == 0 {
		fmt.Println("Array is 32-byte aligned")
	} else {
		fmt.Println("Array is not 32-byte aligned")
	}
	fmt.Printf("Aligned Array Address: %x\n", addr)

	x[0] = 1
	x[1] = 5
	y[0] = 7
	y[1] = 9
	s1, s2, c1, c2 := VecAdd(x, y, carry)
	fmt.Println(s1, s2, c1, c2)
}

/*
func TestVecMul(t *testing.T) {
	// Test case 1: Small numbers
	x1 := []uint64{1, 2}
	y1 := []uint64{4, 5}
	expectedHi0 := uint64(0)
	expectedHi1 := uint64(0)
	expectedLo0 := uint64(4)
	expectedLo1 := uint64(10)
	checkVecMul(t, x1, y1, expectedHi0, expectedHi1, expectedLo0, expectedLo1)

	// Test case 2: Zero values
	x2 := []uint64{0, 0}
	y2 := []uint64{4, 5}
	expectedHi0_2 := uint64(0)
	expectedHi1_2 := uint64(0)
	expectedLo0_2 := uint64(0)
	expectedLo1_2 := uint64(0)
	checkVecMul(t, x2, y2, expectedHi0_2, expectedHi1_2, expectedLo0_2, expectedLo1_2)

	// Test case 3: Multiply by 1
	x3 := []uint64{10, 20}
	y3 := []uint64{1, 1}
	expectedHi0_3 := uint64(0)
	expectedHi1_3 := uint64(0)
	expectedLo0_3 := uint64(10)
	expectedLo1_3 := uint64(20)
	checkVecMul(t, x3, y3, expectedHi0_3, expectedHi1_3, expectedLo0_3, expectedLo1_3)

	// Test case 4: Multiply big integers
	x4 := []uint64{18446744073709551615, 18446744073709551615}
	y4 := []uint64{2, 3}
	expectedHi0_4 := uint64(1)
	expectedHi1_4 := uint64(2)
	expectedLo0_4 := uint64(18446744073709551614)
	expectedLo1_4 := uint64(18446744073709551613)
	checkVecMul(t, x4, y4, expectedHi0_4, expectedHi1_4, expectedLo0_4, expectedLo1_4)

}
*/

func TestVecMulAligned(t *testing.T) {
	// Test case 1: Small numbers
	x := align32Uint64(2)
	y := align32Uint64(2)
	x[0] = 1
	x[1] = 2
	y[0] = 4
	y[1] = 5
	expectedHi0 := uint64(0)
	expectedHi1 := uint64(0)
	expectedLo0 := uint64(4)
	expectedLo1 := uint64(10)
	checkVecMul(t, x, y, expectedHi0, expectedHi1, expectedLo0, expectedLo1)

	// Test case 2: Zero values
	x[0] = 0
	x[1] = 0
	y[0] = 4
	y[1] = 5
	expectedHi0_2 := uint64(0)
	expectedHi1_2 := uint64(0)
	expectedLo0_2 := uint64(0)
	expectedLo1_2 := uint64(0)
	checkVecMul(t, x, y, expectedHi0_2, expectedHi1_2, expectedLo0_2, expectedLo1_2)

	// Test case 3: Multiply by 1
	x[0] = 10
	x[1] = 20
	y[0] = 1
	y[1] = 1
	expectedHi0_3 := uint64(0)
	expectedHi1_3 := uint64(0)
	expectedLo0_3 := uint64(10)
	expectedLo1_3 := uint64(20)
	checkVecMul(t, x, y, expectedHi0_3, expectedHi1_3, expectedLo0_3, expectedLo1_3)

	// Test case 4: Multiply big integers
	x[0] = 18446744073709551615
	x[1] = 18446744073709551615
	y[0] = 2
	y[1] = 3
	expectedHi0_4 := uint64(1)
	expectedHi1_4 := uint64(2)
	expectedLo0_4 := uint64(18446744073709551614)
	expectedLo1_4 := uint64(18446744073709551613)
	checkVecMul(t, x, y, expectedHi0_4, expectedHi1_4, expectedLo0_4, expectedLo1_4)
}

// helper function to check VecMul results against expected values
func checkVecMul(t *testing.T, x, y []uint64, expectedHi0, expectedHi1, expectedLo0, expectedLo1 uint64) {
	hi0, hi1, lo0, lo1 := VecMul(x, y)
	require.Equal(t, expectedHi0, hi0, "hi0 mismatch")
	require.Equal(t, expectedHi1, hi1, "hi1 mismatch")
	require.Equal(t, expectedLo0, lo0, "lo0 mismatch")
	require.Equal(t, expectedLo1, lo1, "lo1 mismatch")
}

func TestCriticalCarry(t *testing.T) {
	hi_p0 := (uint64)(1)
	lo_p0 := (uint64)(2)
	t0 := (uint64)(3)
	d1 := (uint64)(8)

	hi_p1 := (uint64)(1)
	lo_p1 := (uint64)(18446744073709551615)
	t1 := (uint64)(18446744073709551614)
	e1 := (uint64)(3)
	var res1, res2 uint64

	sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei := VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
	_, _, c1, c2 := VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
	res1, res2, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})
	if res1 != 0 && res2 != 3 {
		t.Errorf("Error: expected final carries to be 0 and 2, got %d and %d", res1, res2)
	}

	// Temporary slices for VecAdd_AVX2_I64
	vecX := make([]uint64, 2) // input1
	vecY := make([]uint64, 2) // input2
	vecZ := make([]uint64, 2) // carry/carryOut
	vecU := make([]uint64, 2) // sum

	// sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei = VecAdd([]uint64{t0, t1}, []uint64{d1, e1}, []uint64{0, 0})
	vecX[0], vecX[1] = t0, t1
	vecY[0], vecY[1] = d1, e1
	vecZ[0], vecZ[1] = 0, 0
	VecAdd_AVX2_I64(vecX, vecY, vecZ, vecU)
	sum_t0_di, sum_t1_ei = vecU[0], vecU[1]
	c_t0_di, c_t1_ei = vecZ[0], vecZ[1]
	// d0, e0, c1, c2 = VecAdd([]uint64{lo_p0, lo_p1}, []uint64{sum_t0_di, sum_t1_ei}, []uint64{0, 0})
	vecX[0], vecX[1] = lo_p0, lo_p1
	vecY[0], vecY[1] = sum_t0_di, sum_t1_ei
	vecZ[0], vecZ[1] = 0, 0
	VecAdd_AVX2_I64(vecX, vecY, vecZ, vecU)
	_, _ = vecU[0], vecU[1]
	c1, c2 = vecZ[0], vecZ[1]
	// t0, t1, _, _ = VecAdd([]uint64{hi_p0, hi_p1}, []uint64{c1, c2}, []uint64{c_t0_di, c_t1_ei})
	vecX[0], vecX[1] = hi_p0, hi_p1
	vecY[0], vecY[1] = c1, c2
	vecZ[0], vecZ[1] = c_t0_di, c_t1_ei
	VecAdd_AVX2_I64(vecX, vecY, vecZ, vecU)
	res1, res2 = vecU[0], vecU[1]
	if res1 != 0 && res2 != 3 {
		t.Errorf("Error: expected final carries to be 0 and 2, got %d and %d", res1, res2)
	}
}

func TestVecMontMul(t *testing.T) {
	var x, y, z, expected Element
	val1 := 1
	val2 := 4
	fmt.Println("Setting x to 1")
	x.SetUint64(uint64(val1))

	fmt.Println("Setting y to 4")
	y.SetUint64(uint64(val2))
	expected.SetUint64(uint64(val1 * val2))

	fmt.Println("Multiplying x and y")
	z.Mul(&x, &y)
	if z != expected {
		t.Errorf("Error: expected %d, got %d", expected, z)
	}
}
