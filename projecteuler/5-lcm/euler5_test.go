package main

import (
	"testing"
)

func TestGreatestCommonDenominator(t *testing.T) {
	var out, expected int

	out = gcd(21, 6)
	expected = 3
	if out != expected {
		t.Errorf("GCD incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = gcd(6, 21)
	expected = 3
	if out != expected {
		t.Errorf("GCD incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = gcd(7, 11)
	expected = 1
	if out != expected {
		t.Errorf("GCD incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = gcd(1, 1)
	expected = 1
	if out != expected {
		t.Errorf("GCD incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	var out, expected int

	out = lcm(1, 1)
	expected = 1
	if out != expected {
		t.Errorf("LCM incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = lcm(4, 6)
	expected = 12
	if out != expected {
		t.Errorf("LCM incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = lcm(6, 4)
	expected = 12
	if out != expected {
		t.Errorf("LCM incorrect. Correct: %d. Calculated: %d.", out, expected)
	}

	out = lcm(7, 11) // primes
	expected = 77
	if out != expected {
		t.Errorf("LCM incorrect. Correct: %d. Calculated: %d.", out, expected)
	}
}

func TestProvidedEulerExample(t *testing.T) {
	testInput := 10
	solution := 2520
	result := calculate(testInput)
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}

func TestEulerAnser(t *testing.T) {
	solution := 232792560
	result := calculate(eulerMax)
	if result != solution {
		t.Errorf("Solution was incorrect. Correct: %d. Calculated: %d.", result, solution)
	}
}
