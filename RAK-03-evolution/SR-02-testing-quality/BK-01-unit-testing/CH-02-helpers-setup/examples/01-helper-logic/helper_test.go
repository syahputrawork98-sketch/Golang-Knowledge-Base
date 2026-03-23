package helper

import "testing"

// failureHelper tanpa t.Helper()
func failureHelper(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Error reported inside helper: expected %d, got %d", expected, actual)
	}
}

// betterHelper dengan t.Helper()
func betterHelper(t *testing.T, expected, actual int) {
	t.Helper() // Menandai ini sebagai helper
	if expected != actual {
		t.Errorf("Error reported at caller line: expected %d, got %d", expected, actual)
	}
}

func TestHelperComparison(t *testing.T) {
	// Jalankan sub-test untuk membandingkan output error log
	t.Run("without helper", func(t *testing.T) {
		failureHelper(t, 10, 20)
	})

	t.Run("with helper", func(t *testing.T) {
		betterHelper(t, 10, 20)
	})
}
