package main

import "fmt"

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
	//*slicePtr = (*slicePtr)[0 : len(slice)-1] // Another way of doing this
}

func main() {
	var buffer [100]byte
	slice := buffer[0:50]
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After:  len(slice) =", len(slice))
}
