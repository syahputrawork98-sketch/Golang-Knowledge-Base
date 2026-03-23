package funcinjection

import (
	"errors"
	"testing"
)

func TestGreeterLogic(t *testing.T) {
	t.Run("successful fetch", func(t *testing.T) {
		// Mock fungsional sederhana
		mockFetch := func(id int) (string, error) {
			return "Gopher", nil
		}

		res, err := GreeterLogic(1, mockFetch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != "Hello, Gopher!" {
			t.Errorf("got %s, want Hello, Gopher!", res)
		}
	})

	t.Run("failed fetch", func(t *testing.T) {
		mockFetch := func(id int) (string, error) {
			return "", errors.New("network error")
		}

		_, err := GreeterLogic(1, mockFetch)
		if err == nil {
			t.Fatal("expected error but got none")
		}
	})
}
