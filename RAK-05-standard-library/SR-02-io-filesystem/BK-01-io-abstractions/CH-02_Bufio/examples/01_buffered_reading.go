package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 01_buffered_reading.go - Efisiensi dengan Bufio
// Analogi: Mengambil air dengan ember, bukan dengan gelas satu per satu.

func main() {
	input := "Baris 1: Belajar Go\nBaris 2: Standard Library\nBaris 3: Buffered Reading"
	
	// Membungkus strings.Reader dengan bufio.Scanner
	scanner := bufio.NewScanner(strings.NewReader(input))

	fmt.Println("Scanning lines:")
	count := 1
	for scanner.Scan() {
		// Scan() membaca baris berikutnya secara otomatis
		line := scanner.Text()
		fmt.Printf("[%d] %s\n", count, line)
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
