package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[3] = 100
	a[4] = 200
	fmt.Println("Array a of size 5:", a)
	fmt.Println("get a[3], a[4] value: ", a[3], a[4])

	fmt.Println("Length of array a: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Array twoD: ", twoD)
}