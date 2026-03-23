package main

import (
	"fmt"
	"time"
)

func main() {
	size := 1_000_000

	// Pengujian Tanpa Pre-alokasi (Lambat karena resize memori berulang)
	start := time.Now()
	var s1 []int
	for i := 0; i < size; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("Tanpa Pre-alokasi: %v\n", time.Since(start))

	// Pengujian Dengan Pre-alokasi (Cepat karena memori disiapkan di awal)
	start = time.Now()
	s2 := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("Dengan Pre-alokasi: %v\n", time.Since(start))
}
