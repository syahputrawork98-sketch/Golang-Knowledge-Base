package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{10, 20, 30, 40, 50} // Harus terurut!

	// Binary Search mencari value 40
	idx, found := slices.BinarySearch(nums, 40)
	
	if found {
		fmt.Printf("Data 40 ditemukan di index %d\n", idx)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}

	// Jika kita cari 25 (yang tidak ada)
	idx, found = slices.BinarySearch(nums, 25)
	fmt.Printf("Cari 25 -> Found: %v, Diusulkan di Index: %d (jika ingin diselipkan)\n", found, idx)
}
