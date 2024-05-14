package main
import ("fmt")

func main() {
  s := []int{}
  fmt.Println(len(s))
  fmt.Println(cap(s))
  fmt.Println(s)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}