package main

import "fmt"

func main() {
	fmt.Println("--- Variable Declarations in Go ---")

	// 1. Deklarasi Standar (dengan tipe)
	var name string = "Gopher"

	// 2. Type Inference (compiler menebak tipe)
	var age = 15

	// 3. Short Declaration (Hanya di dalam fungsi)
	location := "Google Cloud"

	fmt.Printf("Name: %s, Age: %d, Location: %s\n", name, age, location)

	// 4. Multiple Declaration
	var x, y, z = 1, 2, "A"
	fmt.Println("Multiple:", x, y, z)
}
