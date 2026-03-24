package main

import (
	"fmt"
	"strings"
)

// 01_search_replace.go - Inspeksi & Manipulasi Teks
// Analogi: Tukang Cetak yang mencari kata tertentu untuk diganti.

func main() {
	text := "Gopher is the official mascot of Go language. Gopher is cute."

	// 1. Pencarian (Search)
	fmt.Printf("Contains 'Go'? %v\n", strings.Contains(text, "Go"))
	fmt.Printf("Has Prefix 'Gopher'? %v\n", strings.HasPrefix(text, "Gopher"))
	fmt.Printf("Index of 'mascot': %d\n", strings.Index(text, "mascot"))

	// 2. Manipulasi (Replace)
	// ReplaceAll membuat salinan baru (string imutabel)
	newText := strings.ReplaceAll(text, "Gopher", "Gopher Engineer")
	
	fmt.Println("\nOriginal:", text)
	fmt.Println("Modified:", newText)

	// 3. Case Transformation
	fmt.Println("Upper:", strings.ToUpper(newText))
}
