/*
Package sumsquares findes the sums of squares and square of sums
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
the calculation without blocking the other calculation.
*/
package sumsquares

func sumOfSquares(sumChan chan int64, max int64) {
	var i int64
	for i = 1; i <= max; i++ {
		sumChan <- i * i
	}
	close(sumChan)
}

func squareOfSums(squareChan chan int64, max int64) {
	var i int64
	sum := int64(0)
	for i = 1; i <= max; i++ {
		sum += i
	}
	squareChan <- sum * sum
	close(squareChan)
}

func squareMinusSum(max int64) (squareMinusSum int64) {
	sumChan := make(chan int64)
	squareChan := make(chan int64)

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
