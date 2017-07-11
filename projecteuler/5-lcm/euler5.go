/*
Finds least common multiples
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
package main

import (
	"fmt"
)

const (
	eulerMax = 20
)

// Greatest common divisor
func gcd(i, j int) int {
	// Euclid's algorithm
	k := i % j
	if k == 0 {
		return j
	}
	return gcd(j, k)
}

// Least common multiple
func lcm(i, j int) int {
	// from LCM wikipedia page
	return i * j / gcd(i, j)
}

func calculate(max int) (lcmOut int) {
	// What is the LCM of all integers between 1 and max, inclusive?
	lcmOut = 1
	for i := 1; i <= max; i++ {
		lcmOut = lcm(lcmOut, i)
	}
	return
}

func main() {
	lcmOut := calculate(eulerMax)
	fmt.Printf("The LCM of numbers from 1 to %d is %d\n", eulerMax, lcmOut)
}
