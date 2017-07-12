package util

import "testing"

var equalsTests = []struct {
	a        []string
	b        []string
	expected bool
}{
	{a: nil, b: nil, expected: true},
	{a: []string{}, b: nil, expected: false},
	{a: []string{"a", "b"}, b: []string{"a", "b"}, expected: true},
	{a: []string{"a", "b"}, b: []string{"a", "b", "c"}, expected: false},
}

func TestEquals(t *testing.T) {
	for _, tt := range equalsTests {
		if Equal(tt.a, tt.b) != tt.expected {
			if tt.expected {
				t.Errorf("incorrectly marked as equal:\na: %v\nb: %v\n", tt.a, tt.b)
			} else {
				t.Errorf("incorrectly marked as unequal:\na: %v\nb: %v\n", tt.a, tt.b)
			}
		}
		// Order of inputs should not matter
		if Equal(tt.b, tt.a) != tt.expected {
			t.Errorf("order of inputs changed answer\na: %v\nb: %v\n", tt.a, tt.b)
		}
	}
}
