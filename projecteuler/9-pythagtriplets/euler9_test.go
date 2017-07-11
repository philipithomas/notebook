package main

import (
	"testing"
)

func TestEulerExample(t *testing.T) {
	sum := 3 + 4 + 5
	out, err := calculatePythagTripletSum(sum)
	expected := 3 * 4 * 5
	if err != nil {
		t.Error("Unexpected error")
	}
	if out != expected {
		t.Errorf("Incorrect answer")
	}
}

func TestTripletExample(t *testing.T) {
	sum := 5 + 12 + 13
	out, err := calculatePythagTripletSum(sum)
	expected := 5 * 12 * 13
	if err != nil {
		t.Error("Unexpected error")
	}
	if out != expected {
		t.Errorf("Incorrect answer")
	}
}

func TestWrongSumThrowsError(t *testing.T) {
	sum := 5 + 12 + 13 + 1
	_, err := calculatePythagTripletSum(sum)
	if err == nil {
		t.Errorf("Error not thrown for incorrect sum %d", sum)
	}
}
