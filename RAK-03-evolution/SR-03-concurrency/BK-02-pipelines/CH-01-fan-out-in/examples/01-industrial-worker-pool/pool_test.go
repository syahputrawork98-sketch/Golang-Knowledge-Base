package workerpool

import (
	"testing"
)

func TestWorkerPool(t *testing.T) {
	const numJobs = 20
	const numWorkers = 5
	
	jobs := make(chan Job, numJobs)
	results := WorkerPool(jobs, numWorkers)

	// Kirim data secara asinkron
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- Job{ID: i, Value: i}
		}
		close(jobs)
	}()

	count := 0
	for range results {
		count++
	}

	if count != numJobs {
		t.Errorf("Expected %d results, got %d", numJobs, count)
	}
}
