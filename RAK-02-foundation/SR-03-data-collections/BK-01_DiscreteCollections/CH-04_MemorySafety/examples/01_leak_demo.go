package main

import (
	"fmt"
	"runtime"
)

func getSmallPart() []int {
	// Membuat array raksasa (10 juta integer)
	giant := make([]int, 10_000_000)
	for i := range giant {
		giant[i] = i
	}
	
	// Kita hanya butuh 3 elemen terakhir
	// TAPI: Ini akan menahan seluruh 10 juta integer di memori!
	return giant[0:3]
}

func main() {
	var m runtime.MemStats
	
	runtime.GC() // Bersihkan dulu
	runtime.ReadMemStats(&m)
	fmt.Printf("Awal - Alloc: %v MB\n", m.Alloc/1024/1024)

	// Ambil bagian kecil
	part := getSmallPart()
	
	runtime.GC() // Coba bersihkan lagi
	runtime.ReadMemStats(&m)
	
	// Hasilnya tetap tinggi karena 'part' menahan 'giant'
	fmt.Printf("Setelah ambil bagian kecil - Alloc: %v MB\n", m.Alloc/1024/1024)
	fmt.Println("Data:", part)
}
