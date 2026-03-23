package main

import "fmt"

func main() {
	fmt.Println("--- Pointer Basics in Go ---")

	x := 42
	ptr := &x // Mengambil alamat memori x

	fmt.Printf("Nilai x: %v\n", x)
	fmt.Printf("Alamat x (&x): %v\n", &x)
	fmt.Printf("Nilai di dalam pointer (ptr): %v\n", ptr)
	fmt.Printf("Nilai di balik pointer (*ptr): %v\n", *ptr)

	// Mengubah nilai asli lewat pointer (Dereferencing)
	*ptr = 100
	fmt.Printf("\nSetelah *ptr = 100, Nilai x menjadi: %v\n", x)

	// Zero value pointer
	var nilPtr *int
	fmt.Printf("Zero value pointer: %v\n", nilPtr)
}
