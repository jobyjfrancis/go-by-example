package main

import (
	"fmt"
)

func square(num int) int {
	return num * num
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Println("The square of ", num, " is: ", square(num))
}
