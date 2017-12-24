package nthprime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEulerExample(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{10001, 104743}, // Euler challenge
	}

	for _, tt := range tests {
		assert.Equal(t, FindNthPrime(tt.n), tt.expected,
			"the %dth prime number should be %d", tt.n, tt.expected)
	}
}
