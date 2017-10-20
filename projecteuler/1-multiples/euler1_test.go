package multiples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	// Test provided example
	assert.Equal(Calculate(10), 23, "The sum with limit 10 should be 23")

	// Test euler problem
	assert.Equal(Calculate(1000), 233168, "The calculated solution should match the correct one")
}
