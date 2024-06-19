package main

import "fmt"

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Create a new instance of the Person struct
	p := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// Access and modify struct fields
	p.Name = "Jane Smith"
	p.Age = 35

	// Print the struct
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:", p.Address)
}