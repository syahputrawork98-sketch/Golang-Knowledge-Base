package main

import "fmt"

type Notifier interface {
	Notify()
}

type Sms struct {
	Phone string
}

// Method 1: Value Receiver
func (s Sms) Notify() {
	fmt.Println("Sending SMS to:", s.Phone)
}

func main() {
	s := Sms{"081234567"}
	
	// Pointer ke Sms memenuhi interface dengan Value Receiver
	var n1 Notifier = &s 
	n1.Notify()

	// Value Sms juga memenuhi (karena Notify() pakai Value Receiver)
	var n2 Notifier = s
	n2.Notify()
	
	fmt.Println("\nKesimpulan: Pointer (*T) adalah 'Warga Negara Kelas Satu' di Inteface.")
	fmt.Println("Ia memenuhi baik fungsionalitas Value maupun Pointer Receiver.")
}
