package main
import "fmt"

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}