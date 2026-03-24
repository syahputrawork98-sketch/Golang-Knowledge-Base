package pool_perf

import (
	"sync"
	"testing"
)

type BigObject struct {
	Data [1024]byte // 1KB Object
}

// 1. Benchmark Tanpa Pool (Alokasi Baru Terus)
func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &BigObject{}
		_ = obj
	}
}

// 2. Setup sync.Pool
var pool = sync.Pool{
	New: func() any {
		return &BigObject{}
	},
}

// 3. Benchmark Dengan Pool (Gunakan Kembali)
func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := pool.Get().(*BigObject)
		// Simulasi penggunaan
		_ = obj
		// Kembalikan ke pool
		pool.Put(obj)
	}
}

/* 
EXECUTION GUIDE:
Jalankan benchmark dengan flag -benchmem untuk melihat perbedaan alokasi:
go test -bench . -benchmem

EXPECTED RESULT:
BenchmarkWithPool akan menunjukkan 0 alokasi per operasi (atau mendekati 0), 
sedangkan BenchmarkNoPool akan menunjukkan alokasi yang konsisten sesuai ukuran struct.
*/
