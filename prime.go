package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Check for divisibility up to the square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	num := 17 // Change this to the number you want to check

	if isPrime(num) {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}
}