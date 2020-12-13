package sorts

import (
	"fmt"
	"math/rand"
	"time"
)

//aqui la que hace el arreglo

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
		fmt.Println(slice[i])
	}
	return slice
}
