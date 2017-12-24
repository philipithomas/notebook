/*
Package sumsquaredifference findes the sums of squares minus the square of sums
13 November 2014

Problem:
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

Solution:
This operation is quite common in statistics as a variance calculation.

I chose to use go routines to calculate the sum and squares concurrently.
Only one value ends up being transmitted on squareChan before it closes,
so it may not be a traditional use of a channel, but it does allow executing
the calculation without blocking the other calculation. In reality, I think
it's more of a fun demonstration of channels rather than the most efficient
approach.
*/
package sumsquaredifference

func sumOfSquares(sumChan chan int, max int) {
	var i int
	for i = 1; i <= max; i++ {
		sumChan <- i * i
	}
	close(sumChan)
}

func squareOfSums(squareChan chan int, max int) {
	var i int
	sum := int(0)
	for i = 1; i <= max; i++ {
		sum += i
	}
	squareChan <- sum * sum
	close(squareChan)
}

// SquareMinusSum returns the sum of squares minus the squre of sums of the
// range of 1 to max
func SquareMinusSum(max int) (squareMinusSum int) {
	sumChan := make(chan int)
	squareChan := make(chan int)

	go squareOfSums(squareChan, max)
	go sumOfSquares(sumChan, max)

	var sumDone, squareDone bool
	for {
		if sumDone && squareDone {
			return
		}
		select {
		case receive, ok := <-sumChan:
			if !ok {
				sumDone = true
			} else {
				squareMinusSum -= receive
			}
		case receive, ok := <-squareChan:
			if !ok {
				squareDone = true
			} else {
				squareMinusSum += receive
			}
		}
	}
}
