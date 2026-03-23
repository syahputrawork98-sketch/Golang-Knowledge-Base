package main

import "fmt"

// Stack kustom yang bisa menampung tipe apa saja (any)
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // Mendapatkan zero value dari tipe T
		return zero, false
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}

func main() {
	// 1. Stack of Integers
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	
	if v, ok := intStack.Pop(); ok {
		fmt.Printf("Popped from intStack: %d (Type: %T)\n", v, v)
	}

	// 2. Stack of Strings
	strStack := Stack[string]{}
	strStack.Push("Go")
	strStack.Push("Generics")

	if v, ok := strStack.Pop(); ok {
		fmt.Printf("Popped from strStack: %s (Type: %T)\n", v, v)
	}
}
