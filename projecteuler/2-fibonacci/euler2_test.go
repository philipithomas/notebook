package fibseq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		lim      int
		expected int
	}{
		{4e6, 4613732},
	}

	for _, tt := range tests {
		assert.Equal(t, Calculate(tt.lim), tt.expected,
			"the sum of even fib numbers below %d is expected to be %d", tt.lim, tt.expected)
	}
}
