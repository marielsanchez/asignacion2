package main

import (
	"fmt"
	"sorts"
)

func main() {

	slice := sorts.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//insertionsort(slice)
	//fmt.Println("\n--- Sorted ---\n\n", slice)

}
