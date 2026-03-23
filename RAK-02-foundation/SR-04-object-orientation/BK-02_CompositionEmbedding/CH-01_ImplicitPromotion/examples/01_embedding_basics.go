package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type Admin struct {
	User  // Embedded field (tanpa nama field)
	Level int
}

func main() {
	a := Admin{
		User: User{
			ID:   1,
			Name: "Antigravity",
		},
		Level: 99,
	}

	// 1. Akses Terpeleset (Promoted)
	fmt.Println("Admin Name:", a.Name) 
	
	// 2. Akses Eksplisit
	fmt.Println("User ID via Admin:", a.User.ID)
	
	fmt.Printf("Admin Struct: %+v\n", a)
}
