package lifecycle

import (
	"fmt"
	"os"
	"testing"
)

// Global state simulasi
var dbConnected bool

func TestMain(m *testing.M) {
	// 1. Setup Global
	fmt.Println(">>> Starting Global Setup (TestMain)...")
	dbConnected = true

	// 2. Jalankan semua test dalam paket
	code := m.Run()

	// 3. Teardown Global
	fmt.Println(">>> Cleaning Up Global resources...")
	dbConnected = false

	// 4. Exit dengan code yang benar
	os.Exit(code)
}

func TestDatabaseRequired(t *testing.T) {
	if !dbConnected {
		t.Fatal("Database should be connected")
	}

	// Local Setup via Cleanup
	t.Cleanup(func() {
		fmt.Println("    - Individual Test Cleanup executed")
	})

	t.Log("Running test with database connection...")
}
