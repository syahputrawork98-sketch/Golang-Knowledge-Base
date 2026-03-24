package broadcast

import (
	"sync"
	"testing"
	"time"
)

func TestCondBroadcast(t *testing.T) {
	var m sync.Mutex
	cond := sync.NewCond(&m)
	ready := false

	const numWorkers = 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// 1. Jalankan para Worker
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			
			m.Lock()
			// PENTING: Selalu gunakan loop for, bukan if, untuk menunggu kondisi
			for !ready {
				t.Logf("Worker %d is waiting...", id)
				cond.Wait() // Menunggu sinyal/broadcast
			}
			t.Logf("Worker %d: Condition met, starting work!", id)
			m.Unlock()
		}(i)
	}

	// 2. Simulasi persiapan oleh Manager
	time.Sleep(1 * time.Second)
	m.Lock()
	t.Log("Manager: Environment is ready. Broadcasting...")
	ready = true
	cond.Broadcast() // Membangunkan SEMUA worker sekaligus
	m.Unlock()

	wg.Wait()
	t.Log("All workers finished.")
}
