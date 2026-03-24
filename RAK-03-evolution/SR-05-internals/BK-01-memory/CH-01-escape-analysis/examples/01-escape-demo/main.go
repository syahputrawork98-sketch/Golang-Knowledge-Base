package main

import "fmt"

type data struct {
	x int
}

// staysOnStack menunjukkan variabel yang tetap di stack
// karena lifetime-nya tidak keluar dari fungsi.
func staysOnStack() int {
	x := 42
	return x // Nilai di-copy, bukan pointer.
}

// escapesToHeap menunjukkan variabel yang "lolos" ke heap
// karena kita mengembalikan pointer ke variabel lokal.
func escapesToHeap() *data {
	d := data{x: 100}
	return &d // Pointer lolos! Dipaksa ke heap.
}

// interfaceEscape menunjukkan alokasi heap karena penggunaan interface
func interfaceEscape(val any) {
	fmt.Println(val)
}

func main() {
	_ = staysOnStack()
	_ = escapesToHeap()
	
	s := "hello heap"
	interfaceEscape(s) // Seringkali memaksa variabel ke heap.
	
	// Lab: Jalankan perintah berikut di terminal:
	// go build -gcflags="-m" main.go
}
