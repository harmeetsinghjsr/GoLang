package main

import "fmt"

func main() {
	sum := 0
	var n int
	fmt.Println("Enter a number n: ")
	fmt.Scanf("%d", &n)
	// for i:=1; i<=n; i++ {
	// 	sum += i
	// }
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("The sum of all even numbers between 1 and n is:", sum)
}