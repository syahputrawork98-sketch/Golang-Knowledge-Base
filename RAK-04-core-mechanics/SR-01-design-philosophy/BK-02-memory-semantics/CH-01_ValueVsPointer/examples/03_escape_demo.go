package main

import "fmt"

// 03_escape_demo.go - Escape Analysis Proof
// Jalankan dengan perintah: go build -gcflags="-m" 03_escape_demo.go

type Item struct {
	ID int
}

func stayOnStack() Item {
	// i dideklarasikan di stack fungsi stayOnStack
	// Karena dikembalikan sebagai value (copy), i tetap aman di stack.
	i := Item{ID: 1}
	return i
}

func escapeToHeap() *Item {
	// j dideklarasikan di stack fungsi escapeToHeap
	// Karena alamatnya (&j) dikembalikan, j "melarikan diri" (escape) ke Heap
	// agar tetap eksis meskipun fungsi ini selesai.
	j := Item{ID: 2}
	return &j 
}

func main() {
	_ = stayOnStack()
	_ = escapeToHeap()
	fmt.Println("Check the build tags output for escape analysis proof!")
}
