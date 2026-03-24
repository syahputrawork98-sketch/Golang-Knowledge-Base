package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func processOrder(ctx context.Context, orderID string) {
	// 1. Mulai Task Utama
	ctx, task := trace.NewTask(ctx, "ProcessOrder")
	defer task.End()

	// Lampirkan data ke task
	trace.Log(ctx, "orderID", orderID)

	// 2. Region: Validation (Single Goroutine)
	region := trace.StartRegion(ctx, "Validation")
	time.Sleep(20 * time.Millisecond) // Simulasi validasi
	region.End()

	// 3. Region: DBUpdate (Menjalankan kerja di Goroutine lain)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Region ini akan diasosiasikan dengan Task induk via context
		r := trace.StartRegion(ctx, "DBUpdate")
		time.Sleep(50 * time.Millisecond)
		r.End()
	}()
	wg.Wait()
}

func main() {
	f, err := os.Create("semantic_trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	log.Println("Starting semantic tracing...")

	ctx := context.Background()
	processOrder(ctx, "ORD-999")

	log.Println("Semantic trace complete. Run:")
	log.Println("  go tool trace semantic_trace.out")
	log.Println("  Look for 'User-defined tasks' in the trace UI.")
}
