package main

import (
	"fmt"
)

const Pi = 3.14

func main() {

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Printf("Type of World variable:%T\n", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth bool = true
	fmt.Println("Go rules?", Truth)

	i := 20
	fmt.Println(i)
}
