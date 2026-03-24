package main

import (
	"fmt"
	"unsafe"
)

// 01_struct_size.go - Membedah Struktur Memori (Alignment & Padding)
// Analogi: Rak buku dengan sekat 8-byte.

type BadStruct struct {
	A bool  // 1 byte
	// Padding 7 bytes
	B int64 // 8 bytes
	C bool  // 1 byte
	// Padding 7 bytes
}

type GoodStruct struct {
	B int64 // 8 bytes
	A bool  // 1 byte
	C bool  // 1 byte
	// Padding 6 bytes
}

func main() {
	bad := BadStruct{}
	good := GoodStruct{}

	fmt.Printf("Ukuran BadStruct: %d bytes\n", unsafe.Sizeof(bad))   // 24 bytes
	fmt.Printf("Ukuran GoodStruct: %d bytes\n", unsafe.Sizeof(good)) // 16 bytes

	fmt.Println("\nOffset BadStruct:")
	fmt.Printf("A (bool): %d\n", unsafe.Offsetof(bad.A))
	fmt.Printf("B (int64): %d\n", unsafe.Offsetof(bad.B)) // Offset 8, ada padding 7 byte di antara A dan B
}
