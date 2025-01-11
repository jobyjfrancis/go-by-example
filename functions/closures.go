package main

import "fmt"

func intSeq() func() int {

	i := 0

	return func() int {
		i++
		return i

	}
}

func main() {

	newInt := intSeq()

	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	newInt2 := intSeq()

	fmt.Println(newInt2())
}
