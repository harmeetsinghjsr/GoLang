package main
import "fmt"

func main() {
  var a int
  var b int
  fmt.Print("Enter two numbers: ")
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Println(a+b)
  fmt.Println(a-b)
  fmt.Println(a*b)
  fmt.Println(a/b)
  fmt.Println(a%b)
}