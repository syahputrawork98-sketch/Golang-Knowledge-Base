package main

import (
	"context"
	"fmt"
	"time"
)

func InfiniteGenerator(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Generator: Sinyal berhenti diterima. Menutup mesin...")
			return
		default:
			fmt.Printf("Generator: Menghasilkan angka %d\n", n)
			n++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Membuat konteks yang bisa dibatalkan secara manual
	ctx, cancel := context.WithCancel(context.Background())

	go InfiniteGenerator(ctx)

	// Biarkan berjalan selama 2 detik
	time.Sleep(2 * time.Second)

	fmt.Println("Main: Memutuskan untuk menghentikan generator sekarang.")
	cancel() // Memanggil fungsi cancel untuk mengirim sinyal Done

	// Tunggu sebentar untuk melihat efek pembersihan
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Aplikasi selesai.")
}
