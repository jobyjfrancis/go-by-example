package main

import "fmt"

func main() {

	s := []string{"Hello", "Welcome", "to", "GoLang"}

	fmt.Println(s)

	fmt.Println("Byte loop:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

}
