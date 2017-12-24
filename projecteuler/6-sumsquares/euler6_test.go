package sumsquaredifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareMinusSum(t *testing.T) {
	// Provided example
	assert.Equal(t, SquareMinusSum(10), 2640)

	// Euler problem
	assert.Equal(t, SquareMinusSum(100), 25164150)
}
