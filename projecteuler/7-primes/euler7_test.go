package main

import (
	"testing"
)

func TestEulerExample(t *testing.T) {
	out := primeNum(6)
	expected := 13
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}

func TestEulerAnswer(t *testing.T) {
	out := primeNum(eulerPrime)
	expected := 104743
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}
