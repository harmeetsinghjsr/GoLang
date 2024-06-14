package main
import "fmt"
func main() {

	var n int
	fmt.Println("Enter number of elements in the array")
	fmt.Scanln(&n)
	arr:=make([]int , n)
	fmt.Println("Enter the elements of the array")
	for i:=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	max:=arr[0]
	for i:=1;i<n;i++{
		if arr[i]>max{
			max=arr[i]
		}
	}
	fmt.Println("The maximum element in the array is",max)
}