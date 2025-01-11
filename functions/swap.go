package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x := "Hello"
	y := "World"
	fmt.Println("String is:", x, y)

	x, y = swap(x, y)
	fmt.Println("After swap:", x, y)
}
