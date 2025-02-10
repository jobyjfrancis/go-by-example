package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	//const s = "Hello"

	fmt.Println("Length of s:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count in string:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U at index %d\n", runeValue, idx)
	}

	fmt.Println("Decode rune in string")

	for i, w := 0, 0; i < len(s); i += w {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", r, i)
		w = size

		examineRune(r)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
