package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("Empty slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("All slice contents: ", s)
	fmt.Println("Slice number 3: ", s[2])

	fmt.Println("Length of the slice: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Updated slice contents: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of slice contents: ", c)

	l := s[2:5]
	fmt.Println("Slice 1: ", l)

	m := s[:5]
	fmt.Println("Slice 2: ", m)

	n := s[2:]
	fmt.Println("Slice 3: ", n)

	t := []string{"g", "h", "i"}
	fmt.Println("Declared slice: ", t)
}
