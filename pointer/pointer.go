package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("Starting value of i:", i)

	zeroval(i)
	fmt.Println("Value of i after call to zeroval function:", i)

	zeroptr(&i)
	fmt.Println("Value of i after call to zeroptr function:", i)
}
