package main

import "fmt"

func greet(name string) func() string {
	n := name

	return func() string {
		msessage := "Hello " + n
		return msessage
	}
}

func main() {
	msg1 := greet("John")
	fmt.Println(msg1())
}
