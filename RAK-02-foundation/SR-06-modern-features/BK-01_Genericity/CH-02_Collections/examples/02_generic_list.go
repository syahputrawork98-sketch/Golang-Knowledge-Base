package main

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type List[T any] struct {
	Head *Node[T]
}

func (l *List[T]) Add(val T) {
	newNode := &Node[T]{Val: val}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *List[T]) Display() {
	for curr := l.Head; curr != nil; curr = curr.Next {
		fmt.Printf("[%v] -> ", curr.Val)
	}
	fmt.Println("NIL")
}

func main() {
	list := List[float64]{}
	list.Add(1.1)
	list.Add(2.2)
	list.Add(3.3)

	fmt.Print("Generic List: ")
	list.Display()
}
