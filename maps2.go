package main
import "fmt"
func main(){
	m := make(map[string]string)
	m["A"]="Apple"
	m["B"]="Ball"
	m["C"]="Cat"
	m["D"]="Dog"
	for key, value := range m {
		fmt.Println(key, value)
	}
	// fmt.Println(m)
}