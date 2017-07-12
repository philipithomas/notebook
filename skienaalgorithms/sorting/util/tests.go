package util

// StringSliceSortTests provides table-driven tests for sorting strings
var StringSliceSortTests = []struct {
	Shuffled []string
	Sorted   []string
}{
	{
		Shuffled: []string{"i", "n", "s", "e", "r", "t", "i", "o", "n", "s", "o", "r", "t"},
		Sorted:   []string{"e", "i", "i", "n", "n", "o", "o", "r", "r", "s", "s", "t", "t"},
	},
	{
		Shuffled: []string{"p", "h", "i", "l", "i", "p"},
		Sorted:   []string{"h", "i", "i", "l", "p", "p"},
	},
}
