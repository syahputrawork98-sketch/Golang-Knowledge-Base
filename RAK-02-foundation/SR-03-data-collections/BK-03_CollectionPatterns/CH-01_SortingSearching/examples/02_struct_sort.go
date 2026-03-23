package main

import (
	"fmt"
	"cmp"    // Paket bantuan pembanding
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Delta", 20},
	}

	// Mengurutkan berdasarkan Umur (Age) - Kecil ke Besar
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Println("Users by Age:", users)

	// Mengurutkan berdasarkan Nama - Besar ke Kecil (Reverse)
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(b.Name, a.Name)
	})
	
	fmt.Println("Users by Name (Desc):", users)
}
