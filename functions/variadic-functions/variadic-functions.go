package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(100, 200)
	sum(1, 2, 3, 4, 5)
	sum(1000)

	nums := []int{10, 20, 30, 40, 50, 60, 70}
	sum(nums...)
}
