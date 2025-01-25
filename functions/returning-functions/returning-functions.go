package main

import "fmt"

func greet(name string) func() {
	return func() {
		fmt.Println("Hello ", name)
	}
}

func main() {
	g1 := greet("John")
	g1()
}
