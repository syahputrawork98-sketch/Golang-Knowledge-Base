package main

import (
	"fmt"
	"slices"
)

func main() {
	// Compact hanya bekerja pada elemen duplikat yang BERURUTAN
	in := []int{1, 1, 2, 3, 3, 3, 4, 1}
	
	// Jika ingin UNIQ total, sebaiknya sort dulu
	slices.Sort(in)
	
	out := slices.Compact(in)

	fmt.Println("Unique Elements:", out)
	fmt.Println("Catatan: Compact sangat efisien karena bekerja in-place.")
}
