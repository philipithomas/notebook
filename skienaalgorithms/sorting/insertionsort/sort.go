// Package insertionsort implements an insertion sort algorithm
package insertionsort

// Sort implements insertionsort for a rune slice
func Sort(item []rune) {
	for i := range item {
		j := i
		for (j > 0) && (item[j] < item[j-1]) {
			item[j], item[j-1] = item[j-1], item[j]
			j--
		}
	}
}
