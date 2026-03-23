package parallel

import (
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	tests := []struct {
		name string
		val  int
	}{
		{"case 1", 1},
		{"case 2", 2},
		{"case 3", 3},
	}

	for _, tt := range tests {
		// TRICKY: Shadowing variabel loop untuk menghindari race condition
		// (Hanya berlaku untuk Go < 1.22, tapi tetap praktik terbaik)
		tt := tt 

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Menjalankan sub-test secara paralel
			time.Sleep(500 * time.Millisecond) // Simulasi kerja berat
			t.Logf("Running case: %s with value %d", tt.name, tt.val)
		})
	}
}
