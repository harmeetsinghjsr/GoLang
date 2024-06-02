package main

import "fmt"

func main() {
	var a, b int = 10, 20

	// Equality operator
	if a == b {
		fmt.Println("a is equal to b")
	}

	// Inequality operator
	if a != b {
		fmt.Println("a is not equal to b")
	}

	// Greater than operator
	if a > b {
		fmt.Println("a is greater than b")
	}

	// Less than operator
	if a < b {
		fmt.Println("a is less than b")
	}

	// Greater than or equal to operator
	if a >= b {
		fmt.Println("a is greater than or equal to b")
	}

	// Less than or equal to operator
	if a <= b {
		fmt.Println("a is less than or equal to b")
	}
}