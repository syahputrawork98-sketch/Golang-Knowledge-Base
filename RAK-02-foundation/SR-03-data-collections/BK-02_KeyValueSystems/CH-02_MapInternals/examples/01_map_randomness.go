package main

import "fmt"

func main() {
	// Map dengan isi yang sama
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}

	fmt.Println("Iterasi 1:")
	for k := range m {
		fmt.Printf("%s ", k)
	}

	fmt.Println("\n\nIterasi 2 (Bisa saja berbeda):")
	for k := range m {
		fmt.Printf("%s ", k)
	}

	fmt.Println("\n\nKesimpulan: Go sengaja memulai iterasi dari bucket acak!")
}
