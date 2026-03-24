package main

import "fmt"

// 02_injected_dependency.go - Decoupling (Testability)
// Analogi: Bohlam lampu dan Fitting. Gampang digonta-ganti.

// Abstraksi
type Database interface {
	Save(data string)
}

// Implementasi Nyata
type MySQL struct{}
func (m MySQL) Save(d string) { fmt.Printf("Saved to MySQL: %s\n", d) }

// Implementasi Mock (Untuk Test)
type MockDB struct{}
func (m MockDB) Save(d string) { fmt.Printf("Logic success: %s\n", d) }

type Service struct {
	db Database // Bergantung pada Abstraksi
}

func main() {
	// Di aplikasi nyata:
	mysql := MySQL{}
	realService := Service{db: mysql}
	realService.db.Save("User Budi")

	// Di unit test:
	mock := MockDB{}
	testService := Service{db: mock}
	testService.db.Save("User Test")
}
