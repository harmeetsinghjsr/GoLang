package main

import (
	"fmt"
)

func search(list []int, target int) bool {
	for _, num := range list {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	var size int
	fmt.Print("Enter the size of the list: ")
	fmt.Scan(&size)

	list := make([]int, size)
	fmt.Println("Enter the elements of the list:")
	for i := 0; i < size; i++ {
		fmt.Scan(&list[i])
	}

	var target int
	fmt.Print("Enter the target number: ")
	fmt.Scan(&target)

	if search(list, target) {
		fmt.Println("Element found!")
	} else {
		fmt.Println("Element not found.")
	}
}