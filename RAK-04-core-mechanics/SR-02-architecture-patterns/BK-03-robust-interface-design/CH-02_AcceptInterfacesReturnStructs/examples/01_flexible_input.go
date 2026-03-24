package main

import "fmt"

// 01_flexible_input.go - Accept Interfaces (Modern API)
// Analogi: Stopkontak universal yang menerima steker jenis apa pun.

type Messenger interface {
	Message() string
}

type SMS struct{}
func (s SMS) Message() string { return "Hi via SMS" }

type Email struct{}
func (e Email) Message() string { return "Hi via Email" }

// ReceiveMessage Menerima Interface (Abstraksi Input)
func ReceiveMessage(m Messenger) {
	fmt.Printf("Message received: %s\n", m.Message())
}

func main() {
	ReceiveMessage(SMS{})
	ReceiveMessage(Email{})
}
