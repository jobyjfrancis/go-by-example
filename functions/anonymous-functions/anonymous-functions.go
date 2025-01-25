package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println(message)
	}("Hello, I am anonymous!!!")
}
