package bigsum

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// Directly check the Euler challenge answer
	assert.Equal(t, Calculate(), "5537376230",
		"the expected first ten digits of the sum are 5537376230")
}

func TestInvalidStringPanics(t *testing.T) {
	assert.Panics(t, func() { sumMod(&big.Int{}, "notanum") },
		"passing this functino an invalid number should panic")
}
