package main

import (
	"testing"
)

const (
	solution = 233168
)

func TestCalculate(t *testing.T) {
	result := calculate()
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}
