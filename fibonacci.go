package main
import "fmt"
func fibo(n int ) int{
	if (n<=1){
		return n
	}
	return fibo(n-1)+fibo(n-2)
}
func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Print(fibo(i)," ")
	}
}