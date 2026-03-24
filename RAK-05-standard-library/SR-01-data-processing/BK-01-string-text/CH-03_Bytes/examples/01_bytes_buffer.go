package main

import (
	"bytes"
	"fmt"
)

// 01_bytes_buffer.go - Membangun Data Biner Dinamis
// Analogi: Papan tulis yang bisa dihapus dan ditulisi ulang.

func main() {
	var buf bytes.Buffer

	// 1. Menulis data ke buffer
	buf.Write([]byte("Header: 0x01, "))
	buf.WriteString("Data: Hello Gophers")
	
	fmt.Printf("Buffer result: %s\n", buf.String())
	fmt.Printf("Buffer length: %d\n", buf.Len())

	// 2. Membaca dari buffer
	p := make([]byte, 8)
	buf.Read(p) // Mengambil 8 byte pertama
	fmt.Printf("First 8 bytes: %s\n", string(p))
	
	// 3. Reset buffer tanpa alokasi ulang memori
	buf.Reset()
	fmt.Printf("After reset, length: %d\n", buf.Len())
}
