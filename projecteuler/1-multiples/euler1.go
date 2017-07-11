/*
Finding multiples of 3 and 5
8 November 2014

Problem: If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.

My approach is fairly straightforward. It uses a named return value sum. Thus, no variable declaration is
needed, and the variable does not need to be explicitly returned. A return must still be called, though.
*/
package main

import (
	"fmt"
)

const (
	limit = 1000
)

func calculate() (sum int) {
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}

func main() {
	sum := calculate()
	fmt.Printf("The sum of multiples of 3 or 5 below %d is %d.\n", limit, sum)
}
