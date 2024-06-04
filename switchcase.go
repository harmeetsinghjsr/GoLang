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
}