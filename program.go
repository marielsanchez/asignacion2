package main

import (
	"sorts"
)

func main() {
	slice := sorts.GenerateSlice(20)
	slice1 := sorts.GenerateSlice(20)
	sorts.HeapSort(slice)
	sorts.InsertionSort(slice1)
}
