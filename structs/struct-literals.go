package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{10, 20}
	v2 = Vertex{X: 120}
	v3 = Vertex{}
	p  = &Vertex{1000, 2000}
)

func main() {
	//v1 := Vertex{10, 20}
	//v2 := Vertex{X: 120}
	//v3 := Vertex{}
	//p := &Vertex{1000, 2000}

	fmt.Println(v1, v2, v3, *p)
}
