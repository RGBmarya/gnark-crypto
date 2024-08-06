package fp

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func TestVecAdd(t *testing.T) {
	// Test case 1: Vectors with zero elements
	in := [8]uint64{}
	// x
	in[0] = 0 // x1
	in[1] = 0 // x2
	// y
	in[2] = 0 // y1
	in[3] = 0 // y2
	//carry
	in[4] = 0
	in[5] = 0
	expectedSum0, expectedCarry0 := bits.Add64(in[0], in[2], in[4])
	expectedSum1, expectedCarry1 := bits.Add64(in[1], in[3], in[5])
	VecAdd_AVX2_I64(&in)
	fmt.Printf("sum1: %d, sum2:%d\n", in[6], in[7])
	fmt.Printf("carry1: %d, sum2:%d\n", in[4], in[5])
	if in[6] != expectedSum0 || in[7] != expectedSum1 || in[4] != expectedCarry0 || in[5] != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			in[6], in[7], in[4], in[5])
	}

	// Test case 2: Vectors with single element
	// x
	in[0] = 1 // x1
	in[1] = 0 // x2
	// y
	in[2] = 2 // y1
	in[3] = 0 // y2
	//carry
	in[4] = 0
	in[5] = 0
	expectedSum0, expectedCarry0 = bits.Add64(in[0], in[2], in[4])
	expectedSum1, expectedCarry1 = bits.Add64(in[1], in[3], in[5])
	VecAdd_AVX2_I64(&in)
	fmt.Printf("sum1: %d, sum2:%d\n", in[6], in[7])
	fmt.Printf("carry1: %d, sum2:%d\n", in[4], in[5])
	if in[6] != expectedSum0 || in[7] != expectedSum1 || in[4] != expectedCarry0 || in[5] != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			in[6], in[7], in[4], in[5])
	}

	// Test case 3: Vectors with multiple elements
	// x
	in[0] = 1 // x1
	in[1] = 2 // x2
	// y
	in[2] = 4 // y1
	in[3] = 5 // y2
	//carry
	in[4] = 0
	in[5] = 0
	expectedSum0, expectedCarry0 = bits.Add64(in[0], in[2], in[4])
	expectedSum1, expectedCarry1 = bits.Add64(in[1], in[3], in[5])
	VecAdd_AVX2_I64(&in)
	fmt.Printf("sum1: %d, sum2:%d\n", in[6], in[7])
	fmt.Printf("carry1: %d, sum2:%d\n", in[4], in[5])
	if in[6] != expectedSum0 || in[7] != expectedSum1 || in[4] != expectedCarry0 || in[5] != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			in[6], in[7], in[4], in[5])
	}

	// Test case 4: Vectors with carry
	// x
	in[0] = 1 // x1
	in[1] = 2 // x2
	// y
	in[2] = 4 // y1
	in[3] = 5 // y2
	//carry
	in[4] = 1
	in[5] = 1
	expectedSum0, expectedCarry0 = bits.Add64(in[0], in[2], in[4])
	expectedSum1, expectedCarry1 = bits.Add64(in[1], in[3], in[5])
	VecAdd_AVX2_I64(&in)
	fmt.Printf("sum1: %d, sum2:%d\n", in[6], in[7])
	fmt.Printf("carry1: %d, sum2:%d\n", in[4], in[5])
	if in[6] != expectedSum0 || in[7] != expectedSum1 || in[4] != expectedCarry0 || in[5] != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			in[6], in[7], in[4], in[5])
	}
	// Test case 5: Vectors with large numbers
	// x
	in[0] = 18446744073709551615 // x1
	in[1] = 18446744073709551615 // x2
	// y
	in[2] = 1 // y1
	in[3] = 1 // y2
	//carry
	in[4] = 0
	in[5] = 0
	expectedSum0, expectedCarry0 = bits.Add64(in[0], in[2], in[4])
	expectedSum1, expectedCarry1 = bits.Add64(in[1], in[3], in[5])
	VecAdd_AVX2_I64(&in)
	fmt.Printf("sum1: %d, sum2:%d\n", in[6], in[7])
	fmt.Printf("carry1: %d, sum2:%d\n", in[4], in[5])
	if in[6] != expectedSum0 || in[7] != expectedSum1 || in[4] != expectedCarry0 || in[5] != expectedCarry1 {
		t.Errorf("Test case 1 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			in[6], in[7], in[4], in[5])
	}
}

/*
func TestVecAddAligned(t *testing.T) {
	x := align32Uint64(2)
	y := align32Uint64(2)
	carry := align32Uint64(2)
	addr := uintptr(unsafe.Pointer(&x[0]))
	if addr%32 == 0 {
		fmt.Println("x Array is 32-byte aligned")
	} else {
		t.Errorf("x Array is not 32-byte aligned")
		return
	}

	fmt.Printf("x Array Address: %x\n", addr)
	addr = uintptr(unsafe.Pointer(&y[0]))
	if addr%32 == 0 {
		fmt.Println("y Array is 32-byte aligned")
	} else {
		t.Errorf("y Array is not 32-byte aligned")
		return
	}
	fmt.Printf("y Array Address: %x\n", addr)

	x[0] = 1
	x[1] = 5
	y[0] = 7
	y[1] = 9
	expectedSum0, expectedCarry0 := bits.Add64(x[0], y[0], 0)
	expectedSum1, expectedCarry1 := bits.Add64(x[1], y[1], 0)
	sum0, sum1, carry0, carry1 := VecAdd(x, y, carry)
	if sum0 != expectedSum0 || sum1 != expectedSum1 || carry0 != expectedCarry0 || carry1 != expectedCarry1 {
		t.Errorf("Test case 5 failed: expected (%d, %d, %d, %d), got (%d, %d, %d, %d)",
			expectedSum0, expectedSum1, expectedCarry0, expectedCarry1,
			sum0, sum1, carry0, carry1)
	}
}
*/

func TestVecMul(t *testing.T) {
	// Test case 1: Small numbers
	x1 := [2]uint64{1, 2}
	y1 := [2]uint64{4, 5}
	expectedHi0 := uint64(0)
	expectedHi1 := uint64(0)
	expectedLo0 := uint64(4)
	expectedLo1 := uint64(10)
	checkVecMul(t, x1, y1, expectedHi0, expectedHi1, expectedLo0, expectedLo1)

	// Test case 2: Zero values
	x2 := [2]uint64{0, 0}
	y2 := [2]uint64{4, 5}
	expectedHi0_2 := uint64(0)
	expectedHi1_2 := uint64(0)
	expectedLo0_2 := uint64(0)
	expectedLo1_2 := uint64(0)
	checkVecMul(t, x2, y2, expectedHi0_2, expectedHi1_2, expectedLo0_2, expectedLo1_2)

	// Test case 3: Multiply by 1
	x3 := [2]uint64{10, 20}
	y3 := [2]uint64{1, 1}
	expectedHi0_3 := uint64(0)
	expectedHi1_3 := uint64(0)
	expectedLo0_3 := uint64(10)
	expectedLo1_3 := uint64(20)
	checkVecMul(t, x3, y3, expectedHi0_3, expectedHi1_3, expectedLo0_3, expectedLo1_3)

	// Test case 4: Multiply big integers
	x4 := [2]uint64{18446744073709551615, 18446744073709551615}
	y4 := [2]uint64{2, 3}
	expectedHi0_4 := uint64(1)
	expectedHi1_4 := uint64(2)
	expectedLo0_4 := uint64(18446744073709551614)
	expectedLo1_4 := uint64(18446744073709551613)
	checkVecMul(t, x4, y4, expectedHi0_4, expectedHi1_4, expectedLo0_4, expectedLo1_4)
}

/*
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
*/

// helper function to check VecMul results against expected values
func checkVecMul(t *testing.T, x, y [2]uint64, expectedHi0, expectedHi1, expectedLo0, expectedLo1 uint64) {
	in := [8]uint64{x[0], x[1], y[0], y[1], 0, 0, 0, 0}
	VecMul_AVX2_I64(&in)
	require.Equal(t, expectedHi0, in[4], "hi0 mismatch")
	require.Equal(t, expectedHi1, in[5], "hi1 mismatch")
	require.Equal(t, expectedLo0, in[6], "lo0 mismatch")
	require.Equal(t, expectedLo1, in[7], "lo1 mismatch")
}

/*
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

	sum_t0_di, sum_t1_ei, c_t0_di, c_t1_ei := VecAdd([2]uint64{t0, t1}, [2]uint64{d1, e1}, [2]uint64{0, 0})
	_, _, c1, c2 := VecAdd([2]uint64{lo_p0, lo_p1}, [2]uint64{sum_t0_di, sum_t1_ei}, [2]uint64{0, 0})
	res1, res2, _, _ = VecAdd([2]uint64{hi_p0, hi_p1}, [2]uint64{c1, c2}, [2]uint64{c_t0_di, c_t1_ei})
	if res1 != 0 && res2 != 3 {
		t.Errorf("Error: expected final carries to be 0 and 2, got %d and %d", res1, res2)
	}

	// Temporary slices for VecAdd_AVX2_I64
	var vecX [2]uint64 // input1
	var vecY [2]uint64 // input2
	var vecZ [2]uint64 // carry/carryOut
	var vecU [2]uint64 // sum

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
*/

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
