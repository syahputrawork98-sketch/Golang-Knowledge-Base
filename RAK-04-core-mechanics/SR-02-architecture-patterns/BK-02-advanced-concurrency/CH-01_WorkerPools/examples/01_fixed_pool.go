package main

import (
	"fmt"
	"time"
)

// 01_fixed_pool.go - Implementasi Worker Pool Sederhana
// Analogi: 3 Teller Bank melayani antrean nasabah.

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: Mulai memproses job %d\n", id, j)
		time.Sleep(time.Second) // Simulasi kerja berat
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Menjalankan 3 worker (Fixed Pool)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Mengirim 5 job ke channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Menandakan tidak ada job lagi

	// Mengumpulkan hasil
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	fmt.Println("Semua job selesai diproses!")
}
