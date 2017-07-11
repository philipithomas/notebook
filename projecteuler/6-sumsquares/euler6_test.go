package main

import (
	"testing"
)

func TestEulerExample(t *testing.T) {
	out := squareMinusSum(int64(10))
	expected := int64(2640)
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}

func TestEulerSolution(t *testing.T) {
	out := squareMinusSum(eulerMax)
	expected := int64(25164150)
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}
