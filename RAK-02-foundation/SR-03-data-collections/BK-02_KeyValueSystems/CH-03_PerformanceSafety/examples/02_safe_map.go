package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	m := make(map[int]int)

	var wg sync.WaitGroup

	// Menulis secara aman dengan Lock
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			m[n] = n
			mu.Unlock()
		}(i)
	}

	// Membaca secara aman dengan RLock (lebih cepat jika banyak pembaca)
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.RLock()
		fmt.Println("Membaca isi map aman...")
		mu.RUnlock()
	}()

	wg.Wait()
	fmt.Println("Selesai. Map aman dari race condition.")
}
