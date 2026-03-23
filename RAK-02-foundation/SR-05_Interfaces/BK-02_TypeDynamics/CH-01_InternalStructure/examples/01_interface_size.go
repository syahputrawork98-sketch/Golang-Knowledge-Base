package main

import (
	"fmt"
	"unsafe"
)

type Empty interface{}
type Reader interface {
	Read()
}

type HeavyStruct struct {
	Data [1000]int
}
func (HeavyStruct) Read() {}

func main() {
	var e Empty
	var r Reader
	var h HeavyStruct

	fmt.Printf("Size of empty interface (any): %d bytes\n", unsafe.Sizeof(e))
	fmt.Printf("Size of non-empty interface:   %d bytes\n", unsafe.Sizeof(r))
	
	fmt.Printf("Size of HeavyStruct:           %d bytes\n", unsafe.Sizeof(h))
	
	// Masukkan struct besar ke interface
	r = h
	fmt.Printf("Size of interface containing HeavyStruct: %d bytes\n", unsafe.Sizeof(r))

	fmt.Println("\nKesimpulan: Ukuran interface SELALU 16 byte (2 word) " +
		"tidak peduli seberapa besar data di dalamnya (karena ia hanya menyimpan pointer).")
}
