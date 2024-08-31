package main
import "fmt"
func main() {
	var i int
	var j int
	var n int = 5
	var choice int
	fmt.Print("Enter choices from 1 to 10: ")
	fmt.Scan(&choice)
	// fmt.Println(" \t")
	switch choice {
	case 1:
		for i = 1; i <= n; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	case 2:
		for i = n; i >= 1; i-- {
			for j = 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	case 3:
		for i = 1; i <= n; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print(i)
			}
			fmt.Println()
		}
	case 4:
		for i = n; i >= 1; i-- {
			for j = 1; j <= i; j++ {
				fmt.Print(i)
			}
			fmt.Println()
		}
	case 5:
		for i = 1; i <= n; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print(j)
			}
			fmt.Println()
		}
	case 6:
		for i = n; i >= 1; i-- {
			for j = 1; j <= i; j++ {
				fmt.Print(j)
			}
			fmt.Println()
		}
	case 7:
		for i = 1; i <= n; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print(n - i + 1)
			}
			fmt.Println()
		}
	case 8:
		for i = n; i >= 1; i-- {
			for j = 1; j <= i; j++ {
				fmt.Print(n - i + 1)
			}
			fmt.Println()
		}
	case 9:
		for i = 1; i <= n; i++ {
			for j = 1; j <= i; j++ {
				fmt.Print(n - j + 1)
			}
			fmt.Println()
		}
	case 10:
		for i = n; i >= 1; i-- {
			for j = 1; j <= i; j++ {
				fmt.Print(n - j + 1)
			}
			fmt.Println()
		}
	default:
		fmt.Println("Invalid choice")
	}
}