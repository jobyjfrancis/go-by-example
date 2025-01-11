// Program to attempt manual calculation of Sqrt of a number

package main

import "fmt"

func Sqrt(x float64) {
	z := 1.0
	for z < 3.0 {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%f", z)
		//z += 1.0

	}
	//return z
}

func main() {
	Sqrt(4)
}
