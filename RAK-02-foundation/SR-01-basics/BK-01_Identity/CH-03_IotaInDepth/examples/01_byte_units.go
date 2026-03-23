package main

import "fmt"

const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << (10*1) = 1024
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
)

func main() {
	fmt.Println("--- Memory Unit Scaling with Iota ---")
	fmt.Printf("1 KB: %d bytes\n", KB)
	fmt.Printf("1 MB: %d bytes\n", MB)
	fmt.Printf("1 GB: %d bytes\n", GB)
	fmt.Printf("1 TB: %d bytes\n", TB)
}
