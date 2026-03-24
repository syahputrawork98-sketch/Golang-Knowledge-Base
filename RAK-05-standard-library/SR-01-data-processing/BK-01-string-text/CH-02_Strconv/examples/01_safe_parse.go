package main

import (
	"fmt"
	"strconv"
)

// 01_safe_parse.go - Konversi Tipe Data yang Aman
// Analogi: Meteran listrik yang memisahkan angka dari teks "kWh".

func main() {
	// 1. String to Int (Atoi)
	strInt := "12345"
	valInt, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Println("Gagal konversi int:", err)
	} else {
		fmt.Printf("Converted Int: %d (Type: %T)\n", valInt, valInt)
	}

	// 2. String to Float
	strFloat := "3.14159"
	valFloat, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Gagal konversi float:", err)
	} else {
		fmt.Printf("Converted Float: %.2f (Type: %T)\n", valFloat, valFloat)
	}

	// 3. String to Bool
	strBool := "true"
	valBool, _ := strconv.ParseBool(strBool)
	fmt.Printf("Converted Bool: %v\n", valBool)
}
