package main

import (
	"fmt"
	"sync"
)

func main() {
	// Eksperimen: Goroutine di dalam loop
	// Sebelum Go 1.22, ini akan mencetak "3, 3, 3" (nilai terakhir i)
	// Sejak Go 1.22, ini secara otomatis mencetak "0, 1, 2" (urutan acak karena goroutine)
	
	var wg sync.WaitGroup
	data := []int{0, 1, 2}

	fmt.Println("Menguji perbaikan Loop Variable di Go 1.22+:")
	for _, v := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Nilai v: %d\n", v)
		}()
	}

	wg.Wait()
	fmt.Println("Selesai. Jika angka di atas unik, maka compiler Anda sudah menggunakan standar 1.22.")
}
