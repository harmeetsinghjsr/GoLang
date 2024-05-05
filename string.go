package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"

	// Length of the string
	length := len(str)
	fmt.Println("Length:", length)

	// Convert to uppercase
	upper := strings.ToUpper(str)
	fmt.Println("Uppercase:", upper)

	// Convert to lowercase
	lower := strings.ToLower(str)
	fmt.Println("Lowercase:", lower)

	// Check if the string contains a substring
	contains := strings.Contains(str, "World")
	fmt.Println("Contains 'World':", contains)

	// Split the string into substrings
	split := strings.Split(str, ",")
	fmt.Println("Split:", split)

	// Join the substrings with a separator
	join := strings.Join(split, "-")
	fmt.Println("Join:", join)
}