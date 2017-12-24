package largestproductinaseries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
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
