package main

import (
	"fmt"
	"sync"
)

// 01_mutex_counter.go - Sinkronisasi State dengan Mutex
// Analogi: Kunci pintu kamar mandi. Hanya satu orang masuk di satu waktu.

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()   // Mengunci akses
	c.v[key]++
	c.mux.Unlock() // Membuka akses
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	wg := sync.WaitGroup{}

	// Menjalankan 100 goroutine untuk menambah counter secara bersamaan
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc("view_count")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", c.v["view_count"])
}
