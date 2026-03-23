package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Initial x (Function Scope):", x)

	if x > 5 {
		// JEBAKAN: Menggunakan := membuat variabel BARU 'x' di dalam blok if
		x := 20 
		fmt.Println("Inside if (Shadowed x):", x)
	}

	// Nilai x di sini TETAP 10, bukan 20!
	fmt.Println("Final x (Function Scope):", x)

	fmt.Println("\n--- Solusi: Gunakan assignment (=) bukan declaration (:=) ---")
	
	if x > 5 {
		x = 100 // Mengupdate variabel x yang ada di luar
	}
	fmt.Println("Correctly updated x:", x)
}
