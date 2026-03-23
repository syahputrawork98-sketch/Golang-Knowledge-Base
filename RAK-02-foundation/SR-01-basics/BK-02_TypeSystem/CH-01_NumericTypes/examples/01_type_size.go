package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 1
	var b int32 = 1
	var c int64 = 1
	var d int = 1 // Mengikuti arsitektur CPU
	var e float64 = 1.0
	var f bool = true

	fmt.Println("--- Go Memory Footprint (in bytes) ---")
	fmt.Printf("int8:    %d byte\n", unsafe.Sizeof(a))
	fmt.Printf("int32:   %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("int64:   %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("int:     %d bytes (pada arsitektur ini)\n", unsafe.Sizeof(d))
	fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(e))
	fmt.Printf("bool:    %d byte\n", unsafe.Sizeof(f))
}
