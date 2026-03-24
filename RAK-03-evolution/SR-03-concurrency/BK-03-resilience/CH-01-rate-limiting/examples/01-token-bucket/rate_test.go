package ratelimit

import (
	"context"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestTokenBucket(t *testing.T) {
	// 1. Inisialisasi Limiter
	// limit: 5 token per detik
	// burst: 3 token (kapasitas ember)
	limiter := rate.NewLimiter(5, 3)

	ctx := context.Background()

	// 2. Simulasi Request Beruntun
	for i := 1; i <= 10; i++ {
		// Wait() akan memblokir goroutine hingga token tersedia
		start := time.Now()
		err := limiter.Wait(ctx)
		if err != nil {
			t.Fatalf("Wait failed: %v", err)
		}
		
		t.Logf("Request %d processed after waiting %v", i, time.Since(start))
	}
}

func TestBurstAllowance(t *testing.T) {
	// Limiter dengan burst 5
	limiter := rate.NewLimiter(1, 5)

	// Burst Pertama: 5 request harus langsung lolos tanpa tunggu
	for i := 1; i <= 5; i++ {
		if !limiter.Allow() {
			t.Errorf("Request %d should have been allowed by burst", i)
		} else {
			t.Logf("Request %d allowed by burst", i)
		}
	}

	// Request ke-6 harus gagal karena burst habis dan refill lambat (1/s)
	if limiter.Allow() {
		t.Error("Request 6 should have been rejected (bucket empty)")
	} else {
		t.Log("Request 6 rejected correctly.")
	}
}
