package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// BadStruct memiliki urutan field yang buruk (boros padding)
type BadStruct struct {
	a bool    // 1 byte
	// 7 bytes padding
	b int64   // 8 bytes
	c bool    // 1 byte
	// 7 bytes padding
}

// GoodStruct memiliki urutan field yang optimal (minim padding)
type GoodStruct struct {
	b int64   // 8 bytes
	a bool    // 1 byte
	c bool    // 1 byte
	// 6 bytes padding
}

func main() {
	bad := BadStruct{}
	good := GoodStruct{}

	fmt.Println("--- Data Alignment & Padding Analysis ---")
	
	fmt.Printf("BadStruct size : %d bytes\n", unsafe.Sizeof(bad))
	fmt.Printf("GoodStruct size: %d bytes\n", unsafe.Sizeof(good))

	fmt.Println("\nOffset Analysis (BadStruct):")
	tBad := reflect.TypeOf(bad)
	for i := 0; i < tBad.NumField(); i++ {
		f := tBad.Field(i)
		fmt.Printf("Field %s: Offset %d, Size %d\n", f.Name, f.Offset, f.Type.Size())
	}

	fmt.Println("\nOffset Analysis (GoodStruct):")
	tGood := reflect.TypeOf(good)
	for i := 0; i < tGood.NumField(); i++ {
		f := tGood.Field(i)
		fmt.Printf("Field %s: Offset %d, Size %d\n", f.Name, f.Offset, f.Type.Size())
	}

	fmt.Println("\nConclusion:")
	fmt.Println("Dengan merapikan field struct dari ukuran TERBESAR ke TERKECIL,")
	fmt.Println("kita bisa menghemat penggunaan RAM secara gratis.")
}
