package fp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVecMul(t *testing.T) {
	// Sample input
	x := []uint64{1, 2}
	y := []uint64{4, 5}

	// Expected output
	expectedHi0 := uint64(0)
	expectedHi1 := uint64(0)
	expectedLo0 := uint64(4)
	expectedLo1 := uint64(10)

	// Call VecMul function
	hi0, hi1, lo0, lo1 := VecMul(x, y)

	// Use require package for assertions
	require.Equal(t, expectedHi0, hi0, "hi0 mismatch")
	require.Equal(t, expectedHi1, hi1, "hi1 mismatch")
	require.Equal(t, expectedLo0, lo0, "lo0 mismatch")
	require.Equal(t, expectedLo1, lo1, "lo1 mismatch")
}
