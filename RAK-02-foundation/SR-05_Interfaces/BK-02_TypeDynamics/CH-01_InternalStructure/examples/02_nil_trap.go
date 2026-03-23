package main

import "fmt"

type Speaker interface {
	Speak()
}

type Human struct{}
func (h *Human) Speak() { fmt.Println("Hello") }

func main() {
	// 1. Nil murni
	var s1 Speaker
	fmt.Println("Is s1 nil?", s1 == nil) // TRUE

	// 2. Typed Nil (Pointer nil yang dimasukkan ke interface)
	var h *Human = nil
	var s2 Speaker = h
	
	fmt.Println("Is h nil?", h == nil)    // TRUE
	fmt.Println("Is s2 nil?", s2 == nil)  // FALSE! (Jebakan!)

	if s2 != nil {
		fmt.Println("s2 is NOT nil because it has Type (*Human), even if Data is nil.")
	}
}
