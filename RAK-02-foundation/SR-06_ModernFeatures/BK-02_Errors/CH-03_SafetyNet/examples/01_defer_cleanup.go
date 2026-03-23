package main

import "fmt"

func ResourceLifecycle() {
	fmt.Println("1. Opening Database Connection")
	defer fmt.Println("5. Closing Database Connection (LIFO - Last)")

	fmt.Println("2. Opening File Stream")
	defer fmt.Println("4. Closing File Stream (LIFO - First)")

	fmt.Println("3. Performing heavy operations...")
}

func main() {
	ResourceLifecycle()
	
	fmt.Println("\nInfo: Perhatikan urutan penutupan (kebalikan dari urutan pembukaan).")
}
