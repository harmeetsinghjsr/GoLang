package main
import ("fmt")
func main() {
	x := 20
	y := 18

	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}
}