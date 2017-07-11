package main

import (
	"testing"
)

func TestProvidedExample(t *testing.T) {
	// Check that we get the right answer for Project Euler's example
	result := calculate(13195)
	solution := 29
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}

func TestCorrectEulerAnswer(t *testing.T) {
	result := calculate(600851475143)
	solution := 6857
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}

func TestPrimeNumbers(t *testing.T) {
	primes := [...]int{2, 3, 5, 7, 11, 149}
	for _, n := range primes {
		if !isPrime(n) {
			t.Errorf("Number %d incorrectly classified as composite", n)
		}
	}

}

func TestCompositeNumbersNotPrime(t *testing.T) {
	composites := [...]int{4, 6, 8, 9, 200, 1024}
	for _, n := range composites {
		if isPrime(n) {
			t.Errorf("Number %d incorrectly classified as prime", n)
		}
	}
}
