package main
import "fmt"

func main(){
	var m map[string]int
	m["a"]=1
	m["e"]=2
	m["i"]=3
	m["o"]=4
	m["u"]=5
	for key, value :=range m{
		fmt.Println(key, value)
	}
}