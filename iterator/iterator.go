package main

import (
	"fmt"
	"iter"
)

type Item int

func getItems() []Item {
	items := []Item{10, 20, 30, 40, 50}
	return items
}

func ItemsIterator(yield func(Item) bool) {
	items := []Item{10, 20, 30, 40, 50}
	for _, v := range items {
		if !yield(v) {
			return
		}
	}
}

func Items() iter.Seq[Item] {
	return func(yield func(Item) bool) {
		items := []Item{10, 20, 30, 40, 50}
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}

}

func main() {
	for i, val := range getItems() {
		fmt.Printf("Items %d: %d\n", i, val)
	}

	for val := range ItemsIterator {
		fmt.Println(val)
		//break
	}

	for val := range Items() {
		fmt.Println(val)
	}
}
