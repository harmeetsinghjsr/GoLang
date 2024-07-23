package main

import "fmt"

func main() {
	// Declare a complex map with string keys and slice values
	complexMap := make(map[string][]int)

	// Add values to the map
	complexMap["key1"] = []int{1, 2, 3}
	complexMap["key2"] = []int{4, 5, 6}
	complexMap["key3"] = []int{7, 8, 9}

	// Access and print values from the map
	fmt.Println(complexMap["key1"]) // Output: [1 2 3]
	fmt.Println(complexMap["key2"]) // Output: [4 5 6]
	fmt.Println(complexMap["key3"]) // Output: [7 8 9]
}