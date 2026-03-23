package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Menggunakan operasi tingkat rendah hardware (Sangat Cepat)
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Hasil Akhir (Atomic): %d\n", atomic.LoadInt64(&count))
}
