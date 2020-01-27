package gcr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreatestCommonDivisor(t *testing.T) {
	var tests = []struct {
		p        uint
		q        uint
		expected uint
	}{
		{0, 0, 0}, // provided example
		{20, 5, 5},
		{5, 20, 5},
		{7, 5, 1}, // prime
		{5, 7, 1},
	}

	for _, tt := range tests {
		actual := GreatestCommonDivisor(tt.p, tt.q)
		assert.Equal(t, tt.expected, actual,
			"expected %d to be the greatest common divisor of %d and %d, but got %d", tt.expected, tt.p, tt.q, actual)
	}
}
