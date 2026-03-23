package main

import "fmt"

func main() {
	// 1. Standard C-Style For
	fmt.Println("--- Standard For ---")
	for i := 0; i < 3; i++ {
		fmt.Println("Iterasi:", i)
	}

	// 2. While-Style For (Hanya kondisi)
	fmt.Println("\n--- While-Style For ---")
	count := 0
	for count < 3 {
		fmt.Println("Counter:", count)
		count++
	}

	// 3. Infinite Loop dengan Break (Do-While style)
	fmt.Println("\n--- Infinite + Break ---")
	x := 0
	for {
		fmt.Println("Running...", x)
		x++
		if x >= 3 {
			break // Menghentikan loop tak terbatas
		}
	}
}
