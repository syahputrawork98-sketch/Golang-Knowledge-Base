package main

import (
	"fmt"
	"runtime"
)

func getSmallPartSafe() []int {
	giant := make([]int, 10_000_000)
	
	// Solusi: Copy data ke slice baru
	res := make([]int, 3)
	copy(res, giant[0:3]) 
	
	return res // 'giant' sekarang bisa dibersihkan oleh GC
}

func main() {
	var m runtime.MemStats
	
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("Awal - Alloc: %v MB\n", m.Alloc/1024/1024)

	part := getSmallPartSafe()
	
	runtime.GC() // GC sekarang bisa menghapus 'giant'
	runtime.ReadMemStats(&m)
	
	// Angka akan kembali rendah!
	fmt.Printf("Setelah aman - Alloc: %v MB\n", m.Alloc/1024/1024)
	fmt.Println("Data:", part)
}
