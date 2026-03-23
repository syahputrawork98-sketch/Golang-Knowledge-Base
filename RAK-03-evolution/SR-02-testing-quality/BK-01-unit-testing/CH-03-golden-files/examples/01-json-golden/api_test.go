package api

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

// Flag -update untuk memperbarui file golden
var update = flag.Bool("update", false, "update golden files")

func TestFetchUserResponse(t *testing.T) {
	actual, err := FetchUserResponse()
	if err != nil {
		t.Fatalf("FetchUserResponse err: %v", err)
	}

	goldenPath := "testdata/response.golden"

	// 1. Logic Update
	if *update {
		if err := os.WriteFile(goldenPath, actual, 0644); err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
	}

	// 2. Logic Comparison
	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	if !bytes.Equal(actual, expected) {
		t.Errorf("Response mismatch!\nActual:\n%s\nExpected:\n%s", actual, expected)
		t.Log("Note: Run 'go test -update' to refresh golden files if change is intentional.")
	}
}
