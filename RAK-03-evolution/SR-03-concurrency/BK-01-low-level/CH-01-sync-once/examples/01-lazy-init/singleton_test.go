package singleton

import (
	"sync"
	"testing"
)

type Database struct {
	ID string
}

var (
	instance *Database
	once     sync.Once
)

// GetInstance mensimulasikan lazy-loading koneksi database yang mahal
func GetInstance() *Database {
	once.Do(func() {
		// Blok ini hanya akan dieksekusi sekali seumur hidup aplikasi
		instance = &Database{ID: "CONNECTED-DB-001"}
	})
	return instance
}

func TestSingleton(t *testing.T) {
	const workers = 50
	var wg sync.WaitGroup
	wg.Add(workers)

	instances := make([]*Database, workers)

	for i := 0; i < workers; i++ {
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}

	wg.Wait()

	// Verifikasi: Semua goroutine harus mendapatkan pointer yang SAMA
	first := instances[0]
	for i, inst := range instances {
		if inst != first {
			t.Errorf("Worker %d got different instance! Singleton failed.", i)
		}
	}
	t.Logf("Success: All %d workers got instance %p", workers, first)
}
