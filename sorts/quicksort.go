package sorts

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func QuickSort(arr []int, low, high int) {
	if low > high {
		return
	}

	pivot := partition(arr, low, high)
	QuickSort(arr, low, pivot-1)
	QuickSort(arr, pivot+1, high)
}
