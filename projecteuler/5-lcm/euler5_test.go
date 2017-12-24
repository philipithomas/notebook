package lcm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	var tests = []struct {
		min      int
		max      int
		expected []int
	}{
		{1, 1, []int{1}},
		{1, 5, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		assert.Equal(t, Range(tt.min, tt.max), tt.expected,
			"the range from %d to %d is expected to be %v", tt.min, tt.max, tt.expected)
	}
}

type gcdTest struct {
	a        int
	b        int
	expected int
}

// nolint: dupl
func TestGreatestCommonDivisor(t *testing.T) {
	var tests = []gcdTest{
		{21, 6, 3},
		{6, 21, 3},
		{7, 11, 1},
		{1, 1, 1},
	}

	for _, tt := range tests {
		assert.Equal(t, GreatestCommonDivisor(tt.a, tt.b), tt.expected,
			"the greatest common divisor of %d and %d should be %d", tt.a, tt.b, tt.expected)
		assert.Equal(t, GreatestCommonDivisor(tt.a, tt.b), GreatestCommonDivisor(tt.b, tt.a),
			"the order of arguments in GreatestCommonDivisor should not matter")
	}
}

type lcmTest struct {
	a        int
	b        int
	expected int
}

// nolint: dupl
func TestLeastCommonMultiple(t *testing.T) {
	var tests = []lcmTest{
		{1, 1, 1},
		{4, 6, 12},
		{6, 4, 12},
		{7, 11, 77},
	}

	for _, tt := range tests {
		assert.Equal(t, LeastCommonMultiple(tt.a, tt.b), tt.expected,
			"the least common multiple of %d and %d should be %d", tt.a, tt.b, tt.expected)
		assert.Equal(t, LeastCommonMultiple(tt.a, tt.b), LeastCommonMultiple(tt.b, tt.a),
			"the order of arguments in LeastCommonMultiple should not matter")
	}

}

func TestLeastCommonMultipleRange(t *testing.T) {
	// Provided euler examplep
	assert.Equal(t, LeastCommonMultipleRange(Range(1, 10)), 2520,
		"the least common multiple of the numbers from 1 to 10 should be 2520")

	// Euler challenge
	assert.Equal(t, LeastCommonMultipleRange(Range(1, 20)), 232792560,
		"the least common multiple of the numbers from 1 to 20 should be 232792560")
}
