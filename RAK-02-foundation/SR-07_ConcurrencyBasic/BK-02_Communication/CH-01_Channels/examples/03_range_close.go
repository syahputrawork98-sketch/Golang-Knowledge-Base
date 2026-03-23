package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	// Produksi data
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * 10
		}
		// SANGAT PENTING: Menutup channel agar range di main berhenti
		close(ch)
	}()

	// Konsumsi data menggunakan range
	fmt.Println("Mulai membaca dari channel:")
	for val := range ch {
		fmt.Printf("Diterima: %d\n", val)
	}

	fmt.Println("Channel telah ditutup, pengulangan selesai.")
}
