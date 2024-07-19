package main

import "fmt"

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Country string
}

// Function to create a new instance of the Person struct
func createPerson(name string, age int, country string) Person {
	return Person{
		Name:    name,
		Age:     age,
		Country: country,
	}
}

// Function to print the details of a person
func printPersonDetails(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Country:", p.Country)
}

// Function to update the age of a person
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

// Function to update the country of a person
func updateCountry(p *Person, newCountry string) {
	p.Country = newCountry
}

func main() {
	// Create a new instance of the Person struct
	p := createPerson("John Doe", 30, "USA")

	// Print the initial details of the person
	printPersonDetails(p)

	// Update the age and country of the person
	updateAge(&p, 31)
	updateCountry(&p, "Canada")

	// Print the updated details of the person
	printPersonDetails(p)
}