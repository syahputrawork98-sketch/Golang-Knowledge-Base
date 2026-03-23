package main

import "fmt"

type Notifier interface {
	Notify()
}

type Email struct {
	Content string
}

// Method Menggunakan Pointer Receiver
func (e *Email) Notify() {
	fmt.Println("Sending Email:", e.Content)
}

func main() {
	e := Email{"Hello Go!"}
	
	// baris di bawah akan ERROR jika diaktifkan (uncomment)
	// var n Notifier = e 
	
	// n.Notify()
	
	fmt.Println("Kesimpulan: Value 'e' tidak memenuhi Notifier karena Notify() pakai pointer.")
	fmt.Println("Error: cannot use e (type Email) as type Notifier in assignment")
}
