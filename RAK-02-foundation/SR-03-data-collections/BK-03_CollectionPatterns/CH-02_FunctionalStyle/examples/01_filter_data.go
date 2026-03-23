package main

import (
	"fmt"
	"slices"
)

func main() {
	scores := []int{10, 55, 30, 80, 5, 100, 42}

	fmt.Println("Original:", scores)

	// Filter: Hapus semua skor di bawah 50
	// Ingat: DeleteFunc mengembalikan slice baru (dengan length baru)
	// tapi tetap menggunakan backing array yang sama.
	scores = slices.DeleteFunc(scores, func(n int) bool {
		return n < 50
	})

	fmt.Println("Filtered (>= 50):", scores)
	fmt.Printf("Len: %d, Cap: %d\n", len(scores), cap(scores))
}
