package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter the day of the week: ")
	fmt.Scanln(&day)

	switch day {
	case 1, 2, 3:
		fmt.Println("It's a weekday")
	case 4, 5:
		fmt.Println("It's almost the weekend")
	case 6, 7:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("that's not a valid day of the week bruh check ur knowledge")
	}
}