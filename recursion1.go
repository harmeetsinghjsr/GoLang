package main
import "fmt"

func testcount(x int) int {
  if x == 16 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}