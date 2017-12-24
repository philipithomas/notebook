package primesieve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		lim      int
		expected int64
	}{
		{10, 17}, // provided example
		{2000000, 142913828922},
	}

	for _, tt := range tests {
		assert.Equal(t, Calculate(tt.lim), tt.expected,
			"the expected sum of prime numbers below %d is %d", tt.lim, tt.expected)
	}
}
