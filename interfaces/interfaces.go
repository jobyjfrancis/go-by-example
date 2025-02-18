package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with radius:", c.radius)
	}
}

func main() {
	r := rectangle{width: 30, height: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
