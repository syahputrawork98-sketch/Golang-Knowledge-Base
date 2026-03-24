package main

import (
	"context"
	"fmt"
	"time"
)

// 01_context_timeout.go - Manajemen Lifecycle (Graceful Cancellation)
// Analogi: Radio prajurit. Jika Jenderal bilang "Waktu habis!", pulang.

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operasi selesai tepat waktu!")
	case <-ctx.Done():
		fmt.Printf("Operasi DIBATALKAN: %v\n", ctx.Err())
	}
}

func main() {
	// Membuat context dengan timeout 1 detik
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Penting untuk menghindari memory leak

	fmt.Println("Main: Memulai operasi lambat (butuh 2 detik)...")
	slowOperation(ctx)
}
