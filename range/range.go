package main

import "fmt"

func main() {

	num := []int{1, 2, 3}
	sum := 0
	for _, i := range num {
		sum += i
	}
	fmt.Println("sum:", sum)

	for i, j := range num {
		if j == 3 {
			fmt.Printf("Index of %d: %d\n", j, i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for i, j := range kvs {
		fmt.Printf("%s --> %s\n", i, j)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "Hello World" {
		fmt.Println(i, c)
	}

}
