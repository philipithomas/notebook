/*
Finds numeric palindromes
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
package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	lower = 100
	upper = 999
)

func isPalindrome(n int) bool {
	// Convert integer to byte array (via string)
	str := []byte(strconv.Itoa(n))
	for i := 0; i < (len(str) / 2); i++ {
		if str[i] != str[(len(str)-i-1)%len(str)] {
			return false
		}
	}
	return true
}

func calculate() (max int, err error) {
	for i := lower; i <= upper; i++ {
		for j := lower; j <= upper; j++ {
			test := i * j
			if isPalindrome(test) {
				if max < test {
					max = test
				}
			}
		}
	}
	if max == 0 {
		err = errors.New("No value found")
	}
	return
}

func main() {
	max, err := calculate()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The max palindrome formed by the multplication of two three-digit numbers is %d\n", max)
}
