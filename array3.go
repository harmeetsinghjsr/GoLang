package main
import "fmt"
func main() {
	var arr = [4]string{"A","B","C","D"}
	fmt.Println(arr)
	arr2:= arr
	fmt.Println("Array_1: ", arr)
    // fmt.Println("Array_2:", arr2)
}