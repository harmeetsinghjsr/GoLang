package main
import "fmt"
func main() {

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	k := 0
	for {
		fmt.Println(k)
		k++
	}

	// Looping through a slice
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Looping through a map
	student := map[string]string{"name": "John", "age ": "20"} 
	for key, value := range student {
		fmt.Println(key, value)
	}
	
	// Looping through a string
	for index, value := range "Hello" {
		fmt.Println(index, string(value))
	}
}