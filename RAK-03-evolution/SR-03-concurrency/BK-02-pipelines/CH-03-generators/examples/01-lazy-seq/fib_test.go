package fib

import (
	"context"
	"testing"
	"time"
)

// FibonacciGenerator menghasilkan deret Fibonacci secara lazy
func FibonacciGenerator(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		a, b := 0, 1
		for {
			select {
			case <-ctx.Done():
				return
			case out <- a:
				a, b = b, a+b
				// Simulasi waktu komputasi
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	return out
}

func TestFibonacciGenerator(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	fibs := FibonacciGenerator(ctx)

	count := 0
	lastVal := -1
	for val := range fibs {
		t.Logf("Fibonacci: %d", val)
		if val < lastVal {
			t.Errorf("Sequence is not increasing: %d < %d", val, lastVal)
		}
		lastVal = val
		count++
	}

	if count == 0 {
		t.Error("Should have generated at least a few numbers")
	}
	t.Logf("Generated %d numbers in 500ms", count)
}
