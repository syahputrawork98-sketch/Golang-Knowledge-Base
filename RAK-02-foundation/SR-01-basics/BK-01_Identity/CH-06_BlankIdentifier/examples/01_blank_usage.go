package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- Blank Identifier Usage ---")

	// 1. Membuang nilai kedua (misalnya error atau index)
	// math.Modf memecah float menjadi bagian bulat dan pecahan
	integer, _ := math.Modf(3.14)
	
	fmt.Printf("Integer part: %v. Fractional part ignored.\n", integer)

	// 2. Loop tanpa butuh index
	items := []string{"Go", "Rust", "C++"}
	for _, item := range items {
		fmt.Printf("Learning %s\n", item)
	}
}
