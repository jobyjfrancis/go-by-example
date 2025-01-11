package main

import (
	"fmt"
)

func sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0

	for _, i := range nums {
		total += i
	}

	fmt.Println("Total:", total)
}

func main() {

	sum(10, 20)
	sum(100, 200, 300, 400, 500, 600, 700)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
