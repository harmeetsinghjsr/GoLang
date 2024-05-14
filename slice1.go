package main
import ("fmt")

func main() {
  sl := make([]int, 5, 10)
  fmt.Printf("sl = %v\n", sl)
  fmt.Printf("length = %d\n", len(sl))
  fmt.Printf("capacity = %d\n", cap(sl))

  // with omitted capacity
  sl12 := make([]int, 5)
  fmt.Printf("sl12 = %v\n", sl12)
  fmt.Printf("length = %d\n", len(sl12))
  fmt.Printf("capacity = %d\n", cap(sl12))
}