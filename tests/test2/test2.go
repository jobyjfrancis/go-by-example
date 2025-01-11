package main

import "fmt"

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	var buffer [100]byte
	slice := buffer[0:50]
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}
