package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ", s, "length: ", len(s), "capacity: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get s[2]: ", s[2])

	fmt.Println("Length of slice, s: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("New slice, s: ", s)
	fmt.Println("Length of slice, s: ", len(s))
	fmt.Println("Capacity of slice, s: ", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Slice c (copied from slice s): ", c)

	l := s[2:5]
	fmt.Println("Slice 1 of s, l: ", l)

	l = s[:5]
	fmt.Println("Slice 2 of s, l: ", l)

	l = s[2:]
	fmt.Println("Slice 3 of s, l: ", l)

	t := []string{"p", "q", "r"}
	fmt.Println("Slice t: ", t)

	t1 := []string{"p", "q", "r"}
	if slices.Equal(t, t1) {
		fmt.Println("t == t1")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
