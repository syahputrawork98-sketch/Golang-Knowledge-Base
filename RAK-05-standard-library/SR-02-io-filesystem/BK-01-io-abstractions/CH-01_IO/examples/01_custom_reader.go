package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// 01_custom_reader.go - Implementasi Kontrak io.Reader
// Analogi: Selang air yang memfilter kotoran (mengubah huruf kecil ke besar).

type UpperReader struct {
	innerReader io.Reader
}

func (r UpperReader) Read(p []byte) (int, error) {
	// 1. Baca data dari reader asli
	n, err := r.innerReader.Read(p)
	if err != nil {
		return n, err
	}

	// 2. Manipulasi data dalam buffer
	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= 32 // Ubah ke UpperCase
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("hello gopher standard library")
	u := UpperReader{innerReader: s}

	var out bytes.Buffer

	// io.Copy memindahkan isi reader ke writer lain.
	if _, err := io.Copy(&out, u); err != nil {
		fmt.Println("copy error:", err)
		return
	}

	fmt.Println("Filtered Output:", out.String())
}
