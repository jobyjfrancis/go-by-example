package main

import "fmt"

func split(num int) (x, y int) {
	x = (num * 4) / 9
	y = num - x
	return
}

func main() {
	n := 17
	fmt.Print("Split of ", n, ":")
	fmt.Println(split(n))
}
