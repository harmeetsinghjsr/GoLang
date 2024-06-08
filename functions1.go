package main
import "fmt"
func sum(a int, b int){
	fmt.Println(a+b)
}
func main() {
	fmt.Println("Enter 2 numbers to add:")
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	// fmt.Println("The sum of", a, "and", b, "is", sum(a, b))
	sum(a, b)
}