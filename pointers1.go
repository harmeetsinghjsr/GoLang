package main

import "fmt"

func changeStrings(pointer1 *string, pointer2 *string) {
	var temp1 string
	var temp2 string
	temp1 = *pointer2
	temp2 = *pointer1
	*pointer1 = temp1
	*pointer2 = temp2
	fmt.Println(*pointer1)
	fmt.Println(*pointer2)
}

func main() {
	var a string = "a"
	var b string = "b"
	changeStrings(&a, &b)
}
