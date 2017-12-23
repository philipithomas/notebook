package main

import (
	"testing"
)

const (
	solution = 4613732
)

func TestCalculate(t *testing.T) {
	result := calculate()
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}
