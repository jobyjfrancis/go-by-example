package main

import "fmt"

func greet(name string) {
	var displayName = func() {
		fmt.Println("Hello ", name)
	}

	displayName()
}

func main() {
	greet("John")
}
