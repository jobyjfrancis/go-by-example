package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	for i, num := range nums {
		if num == 20 {
			fmt.Println("Index of 20: ", i)
		}
	}

}
