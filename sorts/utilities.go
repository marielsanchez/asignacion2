package sorts

import (
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice : max - min
		slice[i] = rand.Intn(999) - 0
	}
	return slice
}

// Generates a slice of size n and in a range of 0 and 31
func GenerateSliceGraph(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice : max - min
		slice[i] = rand.Intn(31) - 0
	}
	return slice
}
