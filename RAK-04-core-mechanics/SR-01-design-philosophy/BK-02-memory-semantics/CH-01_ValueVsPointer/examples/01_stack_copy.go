package main

import (
	"fmt"
)

// 01_stack_copy.go - Value Semantics (Memory Safety)
// Analogi: Fotokopi dokumen. Coretan di salinan tidak mengubah aslinya.

type User struct {
	Name string
	Age  int
}

func updateAge(u User) {
	// u di sini adalah salinan (copy)
	u.Age = 40
	fmt.Printf("Inside updateAge (Copy): %d\n", u.Age)
}

func main() {
	user := User{Name: "Budi", Age: 25}
	
	fmt.Printf("Before updateAge: %d\n", user.Age)
	updateAge(user)
	fmt.Printf("After updateAge: %d\n", user.Age) // Tetap 25
}
