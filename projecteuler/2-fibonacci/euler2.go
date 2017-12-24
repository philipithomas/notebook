/*
Package fibseq calculates the sum of fibonacci sequences' even numbers
9 November 2014

Problem:
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

Solution:
While looping through the calculation would have been about as fast, I chose to try a Go Routine to transmit the Fibonacci numbers.
This allowed me to separate the logic for creating Fibonacci numbers and the logic for computing the sum of even numbers.

One question this raised is how to conduct testing of channels. I have been creating basic unit tests for Project Euler scripts
after I have a confirmed correct answer. This allows me to refactor after finding the correct answer without worry.
*/
package fibseq

const (
	limit = 4e6
)

func emitFib(fibChan chan int, limit int) {

	// Problem says start with [1, 2]. I prefer the more classical [0, 1]
	current := 1
	prior := 0

	// Go doesn't have a "while" - only "for"
	for current <= limit {

		// Emit current
		fibChan <- current

		// Step
		new := current + prior
		current, prior = new, current
	}
	close(fibChan)
}

func processFib(fibChan chan int) (sum int) {
	for {
		val, ok := <-fibChan
		if !ok {
			// closed channel
			return
		}

		// only sum even numbers
		if val%2 == 0 {
			sum += val
		}
	}
}

// Calculate returns the sum of even Fibonacci numbers below a limit
func Calculate(limit int) int {
	fibChan := make(chan int)
	go emitFib(fibChan, limit)

	return processFib(fibChan)
}
