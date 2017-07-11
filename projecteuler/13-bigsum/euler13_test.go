package main

import (
	"testing"
)

const (
	solution = "5537376230"
)

func TestCalculate(t *testing.T) {
	result := calculate()
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %s. Calculated: %s.", solution, result)
	}
}
