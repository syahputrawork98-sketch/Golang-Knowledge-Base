package main

import (
	"context"
	"fmt"
)

type key string

func ProcessRequest(ctx context.Context) {
	// Mengambil value dari konteks
	if requestID := ctx.Value(key("req-id")); requestID != nil {
		fmt.Printf("Worker: Memproses request dengan ID: %v\n", requestID)
	} else {
		fmt.Println("Worker: Request ID tidak ditemukan.")
	}
}

func main() {
	// Menyisipkan metadata ke dalam konteks
	ctx := context.WithValue(context.Background(), key("req-id"), "REQ-12345")

	fmt.Println("Main: Mengirim instruksi ke worker...")
	ProcessRequest(ctx)
}
