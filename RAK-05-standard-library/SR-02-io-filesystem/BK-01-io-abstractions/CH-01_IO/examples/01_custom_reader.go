package main

import (
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

	// Membaca menggunakan io.Copy ke Standard Output
	// io.Copy adalah konsumen universal io.Reader
	fmt.Print("Filtered Output: ")
	io.Copy(nil, u) // Simulasi, biasanya ke os.Stdout
	
	// Lab: Tulis manual ke buffer untuk melihat hasil
	buf := make([]byte, 32)
	n, _ := u.Read(buf)
	fmt.Println(string(buf[:n]))
}
