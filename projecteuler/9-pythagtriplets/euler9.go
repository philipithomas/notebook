/*
Package pythagoreantriplets finds Pythagorean triplets that sum to a
particular number.
20 November 2014

Problem:

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2
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
package pythagoreantriplets

import (
	"errors"
	"math"
)

var (
	// ErrInfeasible is returned when the problem has no solution
	ErrInfeasible = errors.New("infeasible")
)

// Calculate returns the product of a Pythagorean triplet whose digits sum to
// the specified input
func Calculate(sum int) (int, error) { // nolint
	/*
	   See derivation.jpg - but basically, calculate required `b` given
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
	return 0, ErrInfeasible
}
