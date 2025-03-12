package main

import (
	"iterator"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  int
}

func (lst *List[T]) Push(v T) *List[T] {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
	return lst
}

func (lst *List[T]) All() iterator.Seq[T] {
	return func(yield func(T) bool) {

	}
}
