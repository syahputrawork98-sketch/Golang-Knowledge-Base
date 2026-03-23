package main

import (
	"errors"
	"fmt"
)

var ErrDatabaseLost = errors.New("connection lost")

func QueryDatabase() error {
	// Membungkus sentinel error dengan konteks tambahan menggunakan %w
	return fmt.Errorf("failed to load user: %w", ErrDatabaseLost)
}

func main() {
	err := QueryDatabase()
	
	fmt.Println("Full Error String:", err)

	// Mendeteksi error asli di dalam rantai
	if errors.Is(err, ErrDatabaseLost) {
		fmt.Println("Detection: This IS a Database Lost error.")
		fmt.Println("Action: Attempting reconnection...")
	} else {
		fmt.Println("Detection: Unknown error type.")
	}
}
