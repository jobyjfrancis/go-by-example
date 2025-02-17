package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rectangle{width: 10, height: 20}
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perimeter())
}
