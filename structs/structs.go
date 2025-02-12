package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "Alex"
	person1.age = 30
	person1.address = "30A Whitney Street"

	person2.name = "Chris"
	person2.age = 32
	person2.address = "9 Newman Road Rolleston"

	fmt.Println("Person1 Name:", person1.name)
	fmt.Println("Person1 Age:", person1.age)
	fmt.Println("Person1 address:", person1.address)

	fmt.Println("Person2 Name:", person2.name)
	fmt.Println("Person2 Age:", person2.age)
	fmt.Println("Person2 address:", person2.address)
}
