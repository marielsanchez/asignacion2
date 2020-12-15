package sorts

//InsertionSort calls the algoritms to do all the sorting
func InsertionSort(items []int) {
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
