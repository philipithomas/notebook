/*
21 November 2014

Problem:
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.


Solution:

This was remarkably similar to Problem 7, so I decided to reimplement how I answered
the question.

I based my initial solution on a concurrent prime number sieve that Rob Pike mentioned
during a talk on Go Concurrency Patterns at Google I/O 2012. It daisy-chains prime
numbers such that every number emitted from the chain is prime, and it then becomes a
filter. The solution worked, but unfortunately was too slow (~15 minutes) because the
base emitter had no knowledge of discovered primes.

Thus, I rewrote the search algorithm to use a Sieve of Eratosthenes. This algorithm’s
speed surprised me - it calculated the sum of all primes below 2,000,000 in 250ms. The
primary data structure is a slice of booleans.
*/
package main

import (
	"fmt"
)

var (
	eulerLim = 2000000
)

func primeSum(lim int) (sum int64) {
	// False means prime at output
	sieve := make([]bool, lim)

	// Not a prime
	sieve[0] = true
	sieve[1] = true

	prime := 2
	for {
		sum += int64(prime)
		// Get rid of multiples of prime
		for i := prime; i < lim; i = i + prime {
			if i%prime == 0 {
				sieve[i] = true
			}
		}

		// Find next prime
		for {
			prime++
			if prime == lim {
				return
			}
			if !sieve[prime] {
				// found next prime
				break
			}
		}
	}
}

func main() {
	answer := primeSum(eulerLim)
	fmt.Printf("The sum of prime number below %d is %d\n", eulerLim, answer)
}
