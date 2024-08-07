package fp

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
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

// helper function to check VecMul results against expected values
func checkVecMul(t *testing.T, x, y [2]uint64, expectedHi0, expectedHi1, expectedLo0, expectedLo1 uint64) {
	in := [8]uint64{x[0], x[1], y[0], y[1], 0, 0, 0, 0}
	VecMul_AVX2_I64(&in)
	require.Equal(t, expectedHi0, in[4], "hi0 mismatch")
	require.Equal(t, expectedHi1, in[5], "hi1 mismatch")
	require.Equal(t, expectedLo0, in[6], "lo0 mismatch")
	require.Equal(t, expectedLo1, in[7], "lo1 mismatch")
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
