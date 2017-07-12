// Package insertionsort implements an insertion sort algorithm
package insertionsort

// Sort implements insertionsort for a string slice
func Sort(item []string) {
	for i := range item {
		j := i
		for (j > 0) && (item[j] < item[j-1]) {
			item[j], item[j-1] = item[j-1], item[j]
			j = j - 1
		}
	}
}
