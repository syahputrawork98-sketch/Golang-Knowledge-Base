package main

import (
	"errors"
	"fmt"
)

// 02_inspection.go - Inspeksi Error (Advanced Discovery)
// Analogi: Membongkar kotak paket untuk mengecek isi aslinya.

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string { return e.Err.Error() }
func (e *QueryError) Unwrap() error { return e.Err }

func main() {
	// Membuat rantai error: Service -> Repo -> DB
	dbErr := errors.New("timeout di postgres")
	repoErr := &QueryError{Query: "SELECT * FROM users", Err: dbErr}
	serviceErr := fmt.Errorf("service layer failed: %w", repoErr)

	// 1. Pengecekan Eksistensi (Is)
	if errors.Is(serviceErr, dbErr) {
		fmt.Println("Audit: Error disebabkan oleh 'timeout di postgres' (errors.Is detected it)")
	}

	// 2. Ekstraksi Data (As)
	var qErr *QueryError
	if errors.As(serviceErr, &qErr) {
		fmt.Printf("Audit: Query bermasalah adalah '%s' (errors.As extracted it)\n", qErr.Query)
	}
}
