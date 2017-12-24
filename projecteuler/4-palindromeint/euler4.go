/*
Package palindromeint returns the largest palindrome integer of the product of
two numbers in a range.
11 November 2014

Problem:
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Solution:
I put most of my effort into determining whether an integral number is a palindrome
in the isPalindrome function. I used a byte slice data structure to make calculations
on the integral digits more straightforward. Thus, the function compares the first digit
to the last digit, then the second digit to the second-last digit, etc using the byte
slice data structure until meeting in the middle of the slice.
*/
package palindromeint

import (
	"errors"
	"strconv"
)

var (
	// ErrNoPalindromesInRange is returned when no palindromes exist in the provided range
	ErrNoPalindromesInRange = errors.New("no palindromes in the provided range")
)

// IsPalindrome returns whether an integer is the same forwards and backwards.
func IsPalindrome(n int) bool {
	// Convert integer to byte array (via string)
	str := []byte(strconv.Itoa(n))
	for i := 0; i < (len(str) / 2); i++ {
		if str[i] != str[(len(str)-i-1)%len(str)] {
			return false
		}
	}
	return true
}

// LargestPalindromeFromProductOfNumbers returns the largest palindrome integer
// in a range of two multiplied integers, with the bounds of each
// integer as floor to ceil, inclusive to inclusive, of the inputs
func LargestPalindromeFromProductOfNumbers(floor, ceil int) (max int, err error) {
	for i := floor; i <= ceil; i++ {
		for j := floor; j <= ceil; j++ {
			test := i * j
			if IsPalindrome(test) {
				if max < test {
					max = test
				}
			}
		}
	}
	if max == 0 {
		err = ErrNoPalindromesInRange
	}
	return
}
