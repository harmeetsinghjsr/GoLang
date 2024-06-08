package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func addNumbers(a, b int) int {
	return a + b
}

func main() {
	greet("John")
	result := addNumbers(5, 3)
	fmt.Println("Result:", result)
}