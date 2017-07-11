package main

import (
	"testing"
)

func TestEulerExample(t *testing.T) {
	digits := 4
	out := calculateAdjacent(digits)
	expected := int64(5832)
	if out != expected {
		t.Errorf("Calculation of %d adjacent digits incorrect. Correct: %d. Calculated: %d.", digits, out, expected)
	}
}

func TestEulerAnswer(t *testing.T) {
	digits := 13
	out := calculateAdjacent(digits)
	expected := int64(23514624000)
	if out != expected {
		t.Errorf("Calculation of %d adjacent digits incorrect. Correct: %d. Calculated: %d.", digits, out, expected)
	}
}
