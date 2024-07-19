package fp

import (
	"fmt"
	"math/bits"
	"testing"
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
