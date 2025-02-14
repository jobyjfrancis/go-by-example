package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 12
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "John", age: 30})
	fmt.Println(person{name: "Alex"})
	fmt.Println(&person{"Alice", 40})
	fmt.Println(newPerson("James"))

	s := person{name: "Koshy", age: 55}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 57
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
