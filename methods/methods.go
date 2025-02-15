package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectangle{10, 20}
	fmt.Println("Area:", r.area())
	fmt.Println("Area:", r.perimeter())

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Area:", rp.perimeter())
}
