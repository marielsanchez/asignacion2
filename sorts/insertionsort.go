package sorts

import (
	"fmt"
)

//InsertionSort calls the algoritms to do all the sorting
func InsertionSort(numbers []int) {

	fmt.Println("\n--- Unsorted --- \n\n", numbers)
	InsertionSortFunc(numbers)
	fmt.Println("\n--- Sorted by insertion sort---\n\n", numbers)

}

func InsertionSortFunc(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				// swap between j-1 and j
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
