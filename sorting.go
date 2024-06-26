package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define an array of integers
	numbers := []int{5, 2, 9, 1, 7}

	// Sort the array in ascending order
	sort.Ints(numbers)

	// Print the sorted array
	fmt.Println(numbers)
}