package main
import "fmt"
var result int 
func myFunction(x int, y int) int {
  result= x + y
  return result
}
func main() {
  var a, b int
  fmt.Scanf("%d %d", &a, &b)
  fmt.Println(myFunction(a, b))
}