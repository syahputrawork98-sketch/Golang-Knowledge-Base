package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1. Scoped Initialization
	// Variabel 'err' dan 'val' hanya hidup di dalam blok if/else ini.
	// Ini mencegah polusi namespace pada fungsi main.
	if val, err := processData(10); err != nil {
		fmt.Printf("Kesalahan terdeteksi: %v\n", err)
	} else {
		fmt.Printf("Data berhasil diproses: %d\n", val)
	}

	// Mencoba akses 'val' di sini akan menyebabkan error kompilasi:
	// fmt.Println(val) // undefined: val
}

func processData(input int) (int, error) {
	if input < 0 {
		return 0, errors.New("input tidak boleh negatif")
	}
	return input * 2, nil
}
