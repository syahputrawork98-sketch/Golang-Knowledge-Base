package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup

	fmt.Println("Memulai 1000 Goroutines untuk menambah counter...")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // TITIK DATA RACE
		}()
	}

	wg.Wait()
	fmt.Printf("Hasil Akhir: %d\n", count)
	fmt.Println("Catatan: Hasilnya mungkin bukan 1000 karena terjadi Data Race!")
	fmt.Println("Coba jalankan dengan: go run -race 01_race_condition.go")
}
