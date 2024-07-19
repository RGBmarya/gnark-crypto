//go:build !purego
// +build !purego

package bls12377

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/fp"
	"golang.org/x/exp/rand"
)

func TestVecAdd(t *testing.T) {
	rand.Seed(2)
	fp.VecAdd()
}
