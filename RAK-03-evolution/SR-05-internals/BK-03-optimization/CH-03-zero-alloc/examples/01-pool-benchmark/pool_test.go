package pool_demo

import (
	"sync"
	"testing"
)

type LargeStruct struct {
	Data [1024]int
}

// BenchmarkStandardAllocation mensimulasikan pembuatan objek baru terus menerus
func BenchmarkStandardAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &LargeStruct{}
	}
}

var pool = sync.Pool{
	New: func() interface{} {
		return &LargeStruct{}
	},
}

// BenchmarkPoolAllocation mensimulasikan pendaurulangan objek via sync.Pool
func BenchmarkPoolAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := pool.Get().(*LargeStruct)
		// Simulasi penggunaan objek
		obj.Data[0] = 1
		// Kembalikan ke pool
		pool.Put(obj)
	}
}

// Lab Guide:
// Jalankan benchmark di terminal:
// go test -bench . -benchmem
// Perhatikan kolom:
// 1. ns/op (Kecepatan per operasi)
// 2. B/op (Alokasi byte per operasi)
// 3. allocs/op (Jumlah alokasi per operasi)
