package main

import (
	"testing"
)

func TestEulerExample(t *testing.T) {
	out := primeSum(10)
	expected := int64(17)
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d.", expected, out)
	}
}

// Omitted for speed.
func TestEulerAnswer(t *testing.T) {
	out := primeSum(eulerLim)
	expected := int64(142913828922)
	if out != expected {
		t.Errorf("Example incorrect. Correct: %d. Calculated: %d. Diff %d.", expected, out, out-expected)
	}
}
