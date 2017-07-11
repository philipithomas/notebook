/*
Finds pythagorean triplets
20 November 2014

Problem:

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

Solution:

I started off by approaching the problem iteratively - I thought that I
could use the Law of Sines to cleverly figure out which edge to iterate
in order to form a square triangle. However, this proved to be quite difficult
- unit tests were failing, there were tons of edge cases, and I couldn’t
get it to work.

So, I pulled out a pen and paper to work out a possible solution
(see "derivation.jpg" in this package).

I found a way to iterate through a single variable and then determine
the possible other variables. Using this, I implemented a fairly straightforward
programmatic solution.
*/
package main

import (
	"fmt"
	"math"
)

func calculatePythagTripletSum(sum int) (int, error) {
	/*
	   See blog post for derivation - but basically, calculate required `b` given
	   an `a`, then see if a is an integer and if it matches the parameters.
	*/

	// float64 so we can use math standard library
	var a, b float64
	n := float64(sum)

	for b = 1; b < n/2; b++ {
		// Calculate required `a`
		a = (2*b*n - n*n) / (2*b - 2*n)

		// is `a` an integer?
		// (built-in modulus doesn't work on floats)
		if math.Mod(a, 1) == 0 {
			return int(a * b * (n - a - b)), nil
		}
	}
	return 0, fmt.Errorf("Sum %f does not appear to have a triplet", n)
}

func main() {
	eulerSum := 1000
	answer, err := calculatePythagTripletSum(eulerSum)
	if err != nil {
		panic("No answer found")
	}
	fmt.Printf("The product of the digits of a pythagorean triplet of %d is %d \n", eulerSum, answer)
}
