// get all prime factors of the number
package main

import "fmt"
func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	fmt.Print("Prime factors of ", n, " are: ")
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			fmt.Print(i, " ")
			i=i/2
		}
	}
	fmt.Println()
}