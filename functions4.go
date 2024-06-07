package main
import "fmt"
func myFunction(x int, y int) int {	
  return x*y
}
func main() {
  var a, b int
  fmt.Scanf("%d %d", &a, &b)
  result := myFunction(a, b)
  fmt.Println(result)
}
