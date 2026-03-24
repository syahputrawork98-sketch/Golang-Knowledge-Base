package workerpool

import (
	"fmt"
	"sync"
)

// Job mewakili unit kerja
type Job struct {
	ID    int
	Value int
}

// Result mewakili hasil kerja
type Result struct {
	JobID  int
	Output int
}

// WorkerPool mengelola Fan-Out dan Fan-In
func WorkerPool(jobs <-chan Job, numWorkers int) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup

	// FAN-OUT: Menjalankan N worker yang membaca dari channel yang sama
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Simulasi pemrosesan berat
				output := job.Value * 10 
				results <- Result{JobID: job.ID, Output: output}
			}
		}(i)
	}

	// FAN-IN: Menunggu semua worker selesai lalu menutup channel results
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func Example() {
	const numJobs = 10
	jobs := make(chan Job, numJobs)

	// Mulai Pipeline
	results := WorkerPool(jobs, 3)

	// Kirim Pekerjaan
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Value: i}
	}
	close(jobs)

	// Konsumsi Hasil (Sequencing)
	for res := range results {
		fmt.Printf("Job %d completed with output %d\n", res.JobID, res.Output)
	}
}
