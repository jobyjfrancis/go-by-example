package main

import "fmt"

func main() {
	/*	i := 1
		for i <= 3 {
			fmt.Println(i)
			i = i + 1
		}

		fmt.Println()

		for j := 10; j <= 20; j++ {
			fmt.Println(j)
		}

		for {
			fmt.Println("\nHello")
			break
		} */

	for n := 30; n <= 50; n++ {
		if n%2 == 0 {
			fmt.Println(n, " is an even number")
		} else {
			fmt.Println(n, " is an odd number")
		}
	}
}
