package main

import (
	"fmt"
	"os"
)

// 01_file_crud.go - Operasi File Dasar
// Analogi: Asisten Pribadi yang mengelola berkas di laci.

func main() {
	filename := "lab_test.txt"

	// 1. Create & Write
	data := []byte("Gopher Standard Library: OS Track")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}
	fmt.Println("1. File berhasil dibuat.")

	// 2. Read
	content, _ := os.ReadFile(filename)
	fmt.Printf("2. Isi file: %s\n", string(content))

	// 3. Status
	info, _ := os.Stat(filename)
	fmt.Printf("3. Ukuran file: %d bytes\n", info.Size())

	// 4. Delete
	err = os.Remove(filename)
	if err == nil {
		fmt.Println("4. File berhasil dihapus (Cleanup).")
	}
}
