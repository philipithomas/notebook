/*
Package primefactorization calculates the largest prime factor of an integer
10 November 2014

Problem:
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

Solution:
My naive instinct was to start with 1 and check every number until I found the largest number
less than the input that was both prime and a factor. As I continued to ponder, I realized that
the largest possible prime factor of a number is its square root. Thus, by starting with the
square root of the input and decrementing until the test case matched the necessary conditions,
the first match is the correct answer.

I did a better job of unit testing this problem because it involved checking whether an input
was prime, and the project gave an example input and solution.

If I were to improve this solution, I would start by build a better isPrime function.
*/
package primefactorization

import (
	"math"
)

// IsPrime returns whether the given positive integer is prime
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		// include i!=n for formality to pass test for 2
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}

// LargestPrimeFactor returns the largest integral prime factor of a number
func LargestPrimeFactor(x int) int {
	// Start with largest possible prime number
	// "int" will take the floor of the returned float
	n := int(math.Sqrt(float64(x)))

	for {
		// is factor?
		if x%n == 0 {

			if IsPrime(n) {
				// Max prime factor!
				return n
			}
		}
		n--
	}
}
