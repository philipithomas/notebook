package pythagoreantriplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSolutions(t *testing.T) {
	var tests = []struct {
		sum             int
		expectedProduct int
	}{
		{
			sum:             3 + 4 + 5,
			expectedProduct: 3 * 4 * 5,
		},
		{
			sum:             5 + 12 + 13,
			expectedProduct: 5 * 12 * 13,
		},
		{
			// Euler challenge
			sum:             1000,
			expectedProduct: 31875000,
		},
	}

	for _, tt := range tests {
		res, err := Calculate(tt.sum)
		assert.NoError(t, err, "no error is expected")
		assert.Equal(t, res, tt.expectedProduct,
			"digits that sum to %d have an expected product of %d", tt.sum, tt.expectedProduct)
	}
}

func TestNoSolutionReturnsErr(t *testing.T) {
	noSolutionSums := []int{1, 29, 999}
	for _, n := range noSolutionSums {
		_, err := Calculate(n)
		assert.Equal(t, err, ErrInfeasible,
			"no pythagorean triplet expected that sums to %d", n)
	}
}
