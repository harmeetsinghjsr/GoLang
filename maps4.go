package main

import "fmt"

func main() {
	// Creating an empty map
	myMap := make(map[string]int)

	// Adding key-value pairs to the map
	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["cherry"] = 3

	// Accessing values from the map
	fmt.Println("Value of apple:", myMap["apple"])

	// Updating a value in the map
	myMap["banana"] = 5

	// Deleting a key-value pair from the map
	delete(myMap, "cherry")

	// Iterating over the map
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}
}