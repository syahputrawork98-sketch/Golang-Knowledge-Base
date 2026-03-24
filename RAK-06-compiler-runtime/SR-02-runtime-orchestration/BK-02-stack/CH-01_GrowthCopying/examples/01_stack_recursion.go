package main

import (
	"fmt"
)

// 01_stack_recursion.go - Pertumbuhan Stack Dinamis
// Analogi: Koper yang membesar secara otomatis setiap kali Anda menambah baju.

func depthReporter(n int) {
	// Variabel lokal untuk memaksa pemakaian stack
	var dummy [1024]byte
	_ = dummy

	if n % 100 == 0 {
		fmt.Printf("Recursion Depth: %d\n", n)
	}
	
	if n == 0 {
		return
	}
	depthReporter(n - 1)
}

func main() {
	fmt.Println("Starting deep recursion (Go will grow stack automatically)...")
	
	// Go akan melipatgandakan 2KB stack menjadi 4KB, 8KB, dst.
	depthReporter(5000)
	
	fmt.Println("Recursion Finished Safely.")
}
