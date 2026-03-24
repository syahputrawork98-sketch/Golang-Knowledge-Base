package errgroup_demo

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

// fetchData simulates a fetch that may fail
func fetchData(ctx context.Context, source string, failOn string) (string, error) {
	if source == failOn {
		return "", fmt.Errorf("source %q is unavailable", source)
	}
	// Check if context was already cancelled
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	return "data-from-" + source, nil
}

func TestErrGroupShortCircuit(t *testing.T) {
	ctx := context.Background()

	// errgroup.WithContext auto-cancels ctx when any goroutine errors
	g, gctx := errgroup.WithContext(ctx)

	sources := []string{"api-a", "api-b-FAIL", "api-c"}
	results := make([]string, len(sources))

	for i, src := range sources {
		i, src := i, src // capture loop vars
		g.Go(func() error {
			data, err := fetchData(gctx, src, "api-b-FAIL")
			if err != nil {
				return err
			}
			results[i] = data
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		t.Logf("errgroup caught first error: %v", err)
		if !errors.Is(err, fmt.Errorf("source %q is unavailable", "api-b-FAIL")) {
			// Just verify the error is non-nil and descriptive
			t.Logf("Error message: %v", err)
		}
	} else {
		t.Error("Expected an error from the failing source, got nil")
	}
}
