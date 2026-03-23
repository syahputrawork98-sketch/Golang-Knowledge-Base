package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// 1. Definisikan Tabel Uji (Slice of Struct)
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -1, -2},
		{"zero value", 0, 5, 5},
		{"mixed numbers", -1, 5, 4},
	}

	// 2. Iterasi tabel
	for _, tt := range tests {
		// 3. Gunakan t.Run() untuk isolasi kasus
		t.Run(tt.name, func(t *testing.T) {
			actual := Add(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, actual, tt.expected)
			}
		})
	}
}
