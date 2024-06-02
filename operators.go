package main

import "fmt"

func main() {
	// Assignment operators
	var a int = 10
	var b int = 5

	// Logical operators
	fmt.Println(a > b)   // Greater than
	fmt.Println(a < b)   // Less than
	fmt.Println(a == b)  // Equal to
	fmt.Println(a != b)  // Not equal to
	fmt.Println(a >= b)  // Greater than or equal to
	fmt.Println(a <= b)  // Less than or equal to

	// Arithmetic operators
	fmt.Println(a + b)   // Addition
	fmt.Println(a - b)   // Subtraction
	fmt.Println(a * b)   // Multiplication
	fmt.Println(a / b)   // Division
	fmt.Println(a % b)   // Modulus

	// Bitwise operators
	fmt.Println(a & b)   // Bitwise AND
	fmt.Println(a | b)   // Bitwise OR
	fmt.Println(a ^ b)   // Bitwise XOR
	fmt.Println(a << b)  // Left shift
	fmt.Println(a >> b)  // Right shift
	fmt.Println(^a)      // Bitwise complement

	//comparison operators
	fmt.Println(a == b)  // Equal to
	fmt.Println(a != b)  // Not equal to
	fmt.Println(a > b)   // Greater than
	fmt.Println(a < b)   // Less than
	fmt.Println(a >= b)  // Greater than or equal to
	fmt.Println(a <= b)  // Less than or equal to
	
}