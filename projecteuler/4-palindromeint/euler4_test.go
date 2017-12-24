package palindromeint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPalindromeFromProductOfNumbers(t *testing.T) {
	assert := assert.New(t)

	// Instead of testing the "two digit" or "three digit" products explicitly, we
	// test ranges

	// Euler provided example: 2-digit
	ans, err := LargestPalindromeFromProductOfNumbers(10, 99)
	assert.NoError(err)
	assert.Equal(ans, 9009, "the largest palindrome in the range of two-digit numbers product should be 9009")

	// Euler problem: 3-digit
	ans, err = LargestPalindromeFromProductOfNumbers(100, 999)
	assert.NoError(err)
	assert.Equal(ans, 906609, "the largest palindrome in the range of two-digit numbers product should be 9009")

}

func TestInvalidRanges(t *testing.T) {
	assert := assert.New(t)

	var invalidRangeTests = []struct {
		floor int
		ceil  int
	}{
		{10, 10},
		{100, 100},
		{990, 998},
	}
	for _, tt := range invalidRangeTests {
		_, err := LargestPalindromeFromProductOfNumbers(tt.floor, tt.ceil)
		assert.Equal(err, ErrNoPalindromesInRange, "the range %d to %d should have no palindromes", tt.floor, tt.ceil)
	}
}
func TestIntegralPalindromesArePalindromes(t *testing.T) {
	assert := assert.New(t)

	palindromes := [...]int{11, 101, 111, 1221, 12321, 1234321}
	for _, n := range palindromes {
		assert.True(IsPalindrome(n), "%d should be a palindrome")
	}

	notPalindromes := [...]int{10, 12, 13, 113, 201, 311, 301, 1011, 17221, 3141, 12345678}
	for _, n := range notPalindromes {
		assert.False(IsPalindrome(n), "%d should not be a palindrome")
	}

}
