/*
Package lcm finds least common multiples
12 November 2014

Problem:
2520 is the smallest number that can be divided by each
of the numbers from 1 to 10 without any remainder. What is
the smallest positive number that is evenly divisible by
all of the numbers from 1 to 20?

Solution:
The least common multiple (lcm) of two numbers is fairly
straightforward to calculate, and the function is associative
such that lcm(a,b,c) = lcm(a, lcm(a,b)). I used this property
to iteratively calculate the lcm(1,2), then lcm(lcm(1,2),3), etc.
to 20.

This function uses the greatest common divisor (gcd), which is
computed recursively using Euclid’s algorithm.
*/
package lcm

// GreatestCommonDivisor returns the greatest common divisor between two
// positive integers
func GreatestCommonDivisor(i, j int) int {
	// Euclid's algorithm
	k := i % j
	if k == 0 {
		return j
	}
	return GreatestCommonDivisor(j, k)
}

// LeastCommonMultiple returns the least common multiple of two positive
// integers
func LeastCommonMultiple(i, j int) int {
	// from LCM wikipedia page
	return i * j / GreatestCommonDivisor(i, j)
}

// LeastCommonMultipleRange returns the least common multiple of all
// positive integers in the input factors
func LeastCommonMultipleRange(factors []int) (lcm int) {
	lcm = 1
	for _, i := range factors {
		lcm = LeastCommonMultiple(lcm, i)
	}
	return
}

// Range returns an array of integers from min to max, inclusive to inclusive,
// in steps of 1
func Range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
