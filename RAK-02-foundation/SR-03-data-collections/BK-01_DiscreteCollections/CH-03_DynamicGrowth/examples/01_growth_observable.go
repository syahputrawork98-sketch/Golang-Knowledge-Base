package main

import "fmt"

func main() {
	s := make([]int, 0)
	lastCap := cap(s)

	fmt.Printf("Initial - Len: %d, Cap: %d\n", len(s), cap(s))

	// Kita tambahkan 1000 elemen dan amati kapan kapasitas melonjak
	for i := 0; i < 1000; i++ {
		s = append(s, i)
		if cap(s) != lastCap {
			fmt.Printf("Jump! i:%d -> New Cap: %d (Factor: %.2f x)\n", 
				i, cap(s), float64(cap(s))/float64(lastCap))
			lastCap = cap(s)
		}
	}
	
	fmt.Println("\nKesimpulan: Go menggandakan kapasitas di awal, lalu melambat secara halus.")
}
