package main

import (
	"context"
	"fmt"
	"time"
)

func FetchDataFromDB(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Simulasi query berat
		fmt.Println("DB: Data berhasil dibaca.")
	case <-ctx.Done():
		fmt.Printf("DB: Operasi dibatalkan karena: %v\n", ctx.Err())
	}
}

func main() {
	// Membuat konteks dengan batas waktu 2 detik
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Menjamin resource konteks dilepas

	fmt.Println("Main: Memulai Query Database (Batas 2 detik)...")
	FetchDataFromDB(ctx)

	fmt.Println("Main: Melanjutkan ke proses lain.")
}
