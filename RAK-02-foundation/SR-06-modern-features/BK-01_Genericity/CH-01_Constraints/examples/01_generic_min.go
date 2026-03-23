package main

import "fmt"

// Fungsi generic dengan constraint Ordered (simulasi)
// Di sini kita gunakan constraint kustom sederhana
type Number interface {
	int | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 1. Menggunakan Integer
	fmt.Printf("Min Int: %d\n", Min(10, 20))

	// 2. Menggunakan Float64
	fmt.Printf("Min Float: %.2f\n", Min(3.14, 2.71))
	
	fmt.Println("\nKesimpulan: Satu fungsi bekerja untuk banyak tipe tanpa duplikasi kode.")
}
