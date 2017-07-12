package insertionsort

import (
	"testing"

	"github.com/philipithomas/notebook/skienaalgorithms/sorting/util"
)

func TestSort(t *testing.T) {
	for _, tt := range util.StringSliceSortTests {
		Sort(tt.Shuffled)
		if !util.Equal(tt.Shuffled, tt.Sorted) {
			t.Errorf("Sorting incorrect.\nActual: %v\nExpected: %v.", tt.Shuffled, tt.Sorted)
		}
	}
}
