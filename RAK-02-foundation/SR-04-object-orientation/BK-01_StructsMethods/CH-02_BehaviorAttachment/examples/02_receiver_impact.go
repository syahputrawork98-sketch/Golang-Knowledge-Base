package main

import "fmt"

type Counter struct {
	Count int
}

// IncrementValue TIDAK akan mengubah data asli
func (c Counter) IncrementValue() {
	c.Count++
	fmt.Println("Inside IncrementValue:", c.Count)
}

// IncrementPointer AKAN mengubah data asli
func (c *Counter) IncrementPointer() {
	c.Count++
	fmt.Println("Inside IncrementPointer:", c.Count)
}

func main() {
	c := Counter{Count: 0}

	c.IncrementValue()
	fmt.Println("After IncrementValue:", c.Count) // Masih 0

	c.IncrementPointer()
	fmt.Println("After IncrementPointer:", c.Count) // Jadi 1

	fmt.Println("\nKesimpulan: Gunakan Pointer Receiver jika ingin mengubah state " +
		"atau menghindari penyalinan data besar.")
}
