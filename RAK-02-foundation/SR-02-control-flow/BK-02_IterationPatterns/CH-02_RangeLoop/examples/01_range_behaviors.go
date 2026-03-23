package main

import "fmt"

func main() {
	// 1. Range pada Slice (Index & Value)
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("--- Iterasi Slice ---")
	for i, name := range names {
		fmt.Printf("[%d] Halo, %s\n", i, name)
	}

	// 2. Range pada Map (Key & Value - Urutan RANDOM!)
	scores := map[string]int{"Go": 100, "JS": 90, "Python": 95}
	fmt.Println("\n--- Iterasi Map (Urutan tidak terjamin) ---")
	for lang, score := range scores {
		fmt.Printf("Bahasa: %s | Skor: %d\n", lang, score)
	}

	// 3. Range pada String (Unicode Aware)
	text := "Go世界" // '世界' adalah 2 rune, tapi mungkin lebih dari 2 byte
	fmt.Println("\n--- Iterasi String (Rune) ---")
	for pos, char := range text {
		fmt.Printf("Karakter %c di posisi byte %d\n", char, pos)
	}
}
