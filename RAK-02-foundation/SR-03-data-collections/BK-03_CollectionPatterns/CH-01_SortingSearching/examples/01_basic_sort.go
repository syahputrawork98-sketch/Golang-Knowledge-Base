package main

import (
	"fmt"
	"slices" // Go 1.21+
)

func main() {
	// 1. Sortir Angka
	nums := []int{42, 10, 5, 99, 1}
	slices.Sort(nums)
	fmt.Println("Sorted Nums:", nums)

	// 2. Sortir String
	names := []string{"Zebra", "Apple", "Mango", "Banana"}
	slices.Sort(names)
	fmt.Println("Sorted Names:", names)
	
	// 3. Mengecek apakah sudah terurut
	if slices.IsSorted(nums) {
		fmt.Println("Data angka sudah terurut rapi!")
	}
}
