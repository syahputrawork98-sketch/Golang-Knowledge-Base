package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

type Service struct {
	Logger // Embedding Behavior
	Name   string
}

func main() {
	s := Service{Name: "OrderService"}

	// Kita bisa memanggil Log() langsung dari Service
	s.Log("Starting " + s.Name)
	
	fmt.Println("Kesimpulan: Service 'meminjam' behavior dari Logger melalui embedding.")
}
