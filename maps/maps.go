package main

import "fmt"

func main() {

	//var m map[string]int

	//m := make(map[string]int)

	var m = map[string]int{}

	fmt.Println("Initial Map contents: ", m)

	m["k1"] = 100
	m["k2"] = 200

	fmt.Println("Map contents: ", m)
	fmt.Println("Length of map: ", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	delete(m, "k2")
	fmt.Println("Map contents after delete: ", m)

	//_, prs := m["k2"]
	v2, ok := m["k2"]
	fmt.Println("m[k2] is present? ", ok)
	fmt.Println("v2: ", v2)

	n := map[string]int{"foo": 1000, "bar": 2000, "xyz": 3000, "abc": 4000, "fgh": 5000}
	fmt.Println("n: ", n)

	for key, value := range n {
		fmt.Println("Key: ", key, "Value: ", value)
	}

}
