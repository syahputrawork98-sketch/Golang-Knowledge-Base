package race

import (
	"sync"
	"testing"
)

func TestUnsafeCounter(t *testing.T) {
	// 1. Variabel bersama tanpa koordinasi (unsafe)
	counter := 0
	const goroutines = 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	// 2. Akses konkuren tanpa mutex
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			counter++ // Penulisan ke variabel yang sama (RACE!)
		}()
	}

	wg.Wait()

	// 3. Verifikasi: Seringkali nilai tidak mencapai 1000
	t.Logf("Final Counter: %d", counter)
}
