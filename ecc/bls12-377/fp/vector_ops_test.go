package fp

import (
	"fmt"
	"math/bits"
	"testing"

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

func TestVecMul(t *testing.T) {
	// Sample input
	x1 := []uint64{1, 2}
	y1 := []uint64{4, 5}

	// Expected output
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

	// Test case 3: Big Ints
	max := ^uint64(0)
	x3 := []uint64{max, max}
	y3 := []uint64{1, 1}
	expectedHi0_3 := uint64(max)
	expectedHi1_3 := uint64(max)
	expectedLo0_3 := uint64(max)
	expectedLo1_3 := uint64(max)
	checkVecMul(t, x3, y3, expectedHi0_3, expectedHi1_3, expectedLo0_3, expectedLo1_3)

}

// Helper function to check VecMul results against expected values
func checkVecMul(t *testing.T, x, y []uint64, expectedHi0, expectedHi1, expectedLo0, expectedLo1 uint64) {
	hi0, hi1, lo0, lo1 := VecMul(x, y)
	require.Equal(t, expectedHi0, hi0, "hi0 mismatch")
	require.Equal(t, expectedHi1, hi1, "hi1 mismatch")
	require.Equal(t, expectedLo0, lo0, "lo0 mismatch")
	require.Equal(t, expectedLo1, lo1, "lo1 mismatch")
}
