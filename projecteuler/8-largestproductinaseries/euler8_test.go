package largestproductinaseries

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAdjacentProductSuccess(t *testing.T) {
	var tests = []struct {
		numAdjacent int
		expected    int64
	}{
		{1, 9},
		{4, 5832},         // provided example
		{13, 23514624000}, // euler challenge
	}
	for _, tt := range tests {
		assert.Equal(t, LargestAdjacentProduct(tt.numAdjacent), tt.expected,
			"the largest product of %d adjacent numbers should be %d", tt.numAdjacent, tt.expected)
	}
}

func TestPanicOnError(t *testing.T) {
	assert.Panics(t, func() { panicOnError(errors.New("error")) }, "expected the code to panic")
	assert.NotPanics(t, func() { panicOnError(nil) }, "no error should not trigger a panic")
}
