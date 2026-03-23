package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         // Kunci akses
	defer c.mu.Unlock() // Pastikan buka kunci di akhir
	c.value++
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Hasil Akhir (Safe): %d\n", counter.value)
}
