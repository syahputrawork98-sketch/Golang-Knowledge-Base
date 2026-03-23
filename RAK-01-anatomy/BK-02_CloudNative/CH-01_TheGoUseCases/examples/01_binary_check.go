package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("--- Cloud Native Portability Test ---")
	fmt.Printf("Program ini dikompilasi untuk: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("Meskipun dijalankan di server tanpa Go, saya akan tetap berjalan.")
	fmt.Println("Gunakan perintah 'go build -o test_binary 01_binary_check.go' untuk melihat hasilnya.")
}
