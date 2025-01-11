package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{100, 200}
	fmt.Println("Origin values:", v)
	v.X = 120
	v.Y = 230
	fmt.Println("After changing:", v)
}
