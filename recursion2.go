package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Print("Enter a number: ")
  var n float64
  fmt.Scan(&n)
  fmt.Println("Factorial of",n,"is",factorial_recursion(n))
}