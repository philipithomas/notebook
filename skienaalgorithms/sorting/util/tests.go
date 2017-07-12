package util

// RuneSliceSortTests provides table-driven tests for sorting strings
var RuneSliceSortTests = []struct {
	Shuffled []rune
	Sorted   []rune
}{
	{
		Shuffled: []rune{'i', 'n', 's', 'e', 'r', 't', 'i', 'o', 'n', 's', 'o', 'r', 't'},
		Sorted:   []rune{'e', 'i', 'i', 'n', 'n', 'o', 'o', 'r', 'r', 's', 's', 't', 't'},
	},
	{
		Shuffled: []rune{'p', 'h', 'i', 'l', 'i', 'p'},
		Sorted:   []rune{'h', 'i', 'i', 'l', 'p', 'p'},
	},
}
