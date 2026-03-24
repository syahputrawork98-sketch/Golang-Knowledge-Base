package singleflight_demo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"golang.org/x/sync/singleflight"
)

// dbCallCount tracks how many times the "database" is actually hit
var dbCallCount int64

// fetchFromDB simulates an expensive database operation
func fetchFromDB(userID string) (string, error) {
	atomic.AddInt64(&dbCallCount, 1)
	return fmt.Sprintf("user-data-for-%s", userID), nil
}

func TestSingleflight_NoDuplicateCalls(t *testing.T) {
	var g singleflight.Group
	dbCallCount = 0

	const numGoroutines = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	results := make([]string, numGoroutines)

	// 50 goroutines all request the same user simultaneously
	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()

			// Only ONE of these will actually call fetchFromDB.
			// The other 49 will receive the same result.
			v, err, shared := g.Do("user-1", func() (interface{}, error) {
				return fetchFromDB("user-1")
			})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			results[idx] = v.(string)
			if shared {
				t.Logf("Goroutine %d: Received shared result", idx)
			}
		}(i)
	}

	wg.Wait()

	// Key assertion: DB should have been called only ONCE
	if dbCallCount != 1 {
		t.Errorf("Expected DB to be called 1 time, but was called %d times", dbCallCount)
	}
	t.Logf("SUCCESS: %d goroutines made requests, DB was called exactly %d time(s)", numGoroutines, dbCallCount)

	// Verify all results are consistent
	for i, r := range results {
		if r != "user-data-for-user-1" {
			t.Errorf("Goroutine %d got unexpected result: %s", i, r)
		}
	}
}
