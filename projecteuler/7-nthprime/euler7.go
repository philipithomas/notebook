/*
Package nthprime finds the nth prime number
14 November 2014

Problem:
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10,001st prime number?

Solution:
The naive way to approach this would be to go through every number,
then check that number to see if it had any factors, until I found
10,001 such numbers. However, my solution to problem 5 reminded me
that any non-prime number may be factored into primes, and by caching
every prime number I found, I could check the modulus of subsequent
numbers against the cache.

Thus, to find the 10,001st prime number, I only need to divide candidates
by 10,000 numbers. This sounds high - but it is actually quite efficient!
In fact, finding 10,001 primes using this script on my Macbook Air took 541ms.
That’s fast! My method for identifying a prime number is efficient, but
perhaps there may be a better way to search for prime number candidates.
*/
package nthprime

// FindNthPrime returns the Nth prime number
func FindNthPrime(n int) int {
	// Start with known 2 as prime
	i := 2
	primes := []int{i}
	for {
		if len(primes) == n {
			return i
		}

		i++
		isPrime := true
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}
}
