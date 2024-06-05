package main
import "fmt"
func Author(book string, author string, year int){
	fmt.Println("The author of", book, "is", author, "and it was published in", year)
}
func main() {
	fmt.Println("Enter the name of a book, the author and the year it was published:")
	var book, author string
	var year int
	fmt.Scanf("%s %s %d", &book, &author, &year)
	Author(book, author, year)
}