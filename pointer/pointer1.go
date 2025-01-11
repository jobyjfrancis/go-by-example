package main

import "fmt"

func main() {
	i, j := 120, 130
	p := &i
	fmt.Println("Address of variable i: ", p)
	fmt.Println("Initial value of variable i: ", *p)
	*p = 150
	fmt.Println("Current value of i:", i)

	p = &j
	fmt.Println("\nAddress of variable j: ", p)
	fmt.Println("Initial value of variable j:", *p)
	*p = *p / 5
	fmt.Println("Current value of j:", j)
}
