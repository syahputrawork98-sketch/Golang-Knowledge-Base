package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Pastikan Done dipanggil di akhir, apa pun yang terjadi
	defer wg.Done()

	fmt.Printf("Worker %d: Starting task...\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d: Task Completed!\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 5
	fmt.Printf("Orchestrating %d workers...\n\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// 1. Tambah counter SEBELUM meluncurkan goroutine
		wg.Add(1)
		go worker(i, &wg) // Kirim sebagai POINTER
	}

	// 2. Blokir main sampai counter kembali ke 0
	fmt.Println("Main: Waiting for all workers to finish...")
	wg.Wait()
	
	fmt.Println("\nMain: All workers finished. Shutting down system.")
}
