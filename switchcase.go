package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Print("Enter a command: ")
	fmt.Scanln(&n)

	switch n {
	case "start":
		fmt.Println("Starting the process...")
	case "stop":
		fmt.Println("Stopping the process...")
	case "restart":
		fmt.Println("Restarting the process...")
	default:
		fmt.Println("Invalid command!")
	}
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&a)
	switch a {
	case 1:
		fmt.Println("a is equal to 1")
	case 2:
		fmt.Println("a is equal to 2")
	case 3:
		fmt.Println("a is equal to 3")
	default:
		fmt.Println("a is not equal to 1, 2, or 3")

}
}