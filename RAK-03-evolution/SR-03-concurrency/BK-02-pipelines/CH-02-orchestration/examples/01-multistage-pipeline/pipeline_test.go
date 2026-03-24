package orchestration

import (
	"context"
	"testing"
	"time"
)

// Stage 1: Generator (Source)
func Generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

// Stage 2: Square (Processor)
func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

func TestPipelineOrchestration(t *testing.T) {
	// Setup Context dengan Timeout untuk simulasi pembatalan
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Building Pipeline
	nums := []int{1, 2, 3, 4, 5}
	stage1 := Generator(ctx, nums...)
	stage2 := Square(ctx, stage1)

	// Consuming Results
	var finalResults []int
	for res := range stage2 {
		t.Logf("Processed: %d", res)
		finalResults = append(finalResults, res)
	}

	if len(finalResults) != len(nums) {
		t.Errorf("Expected %d results, got %d", len(nums), len(finalResults))
	}
}

func TestPipelineCancellation(t *testing.T) {
	// Setup Context yang langsung dibatalkan
	ctx, cancel := context.WithCancel(context.Background())
	
	stage1 := Generator(ctx, 1, 2, 3, 4, 5)
	cancel() // Batalkan segera

	// Pipeline harus berhenti dengan cepat tanpa deadlock
	_, ok := <-stage1
	if ok {
		t.Error("Pipeline should have closed immediately after cancellation")
	}
}
