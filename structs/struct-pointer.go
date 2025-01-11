package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{10, 20}
	p := &v
	fmt.Println(*p)
	fmt.Println(&v)
	fmt.Println(&v.X)
	fmt.Println()

	p.X = 10e9
	fmt.Println(v)
	fmt.Println()

	// Testing a normal integer variable and its pointer
	var x int = 120
	p1 := &x
	fmt.Println(p1)
	fmt.Println(*p1)
}
