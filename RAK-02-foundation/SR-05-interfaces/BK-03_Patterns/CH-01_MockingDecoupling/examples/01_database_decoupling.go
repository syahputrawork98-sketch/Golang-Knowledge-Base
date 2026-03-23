package main

import "fmt"

// Interface didefinisikan oleh LOGIKA BISNIS
type DataStore interface {
	GetUser(id int) string
}

// Logika Bisnis (Service)
type UserService struct {
	store DataStore // Bergantung pada Abstraksi
}

func (u UserService) WelcomeUser(id int) {
	name := u.store.GetUser(id)
	fmt.Printf("Welcome to the system, %s!\n", name)
}

// Implementasi SQL (Low Level)
type SQLStore struct{}
func (SQLStore) GetUser(id int) string {
	return fmt.Sprintf("UserFromSQL#%d", id)
}

func main() {
	// Menghubungkan komponen (Wiring)
	db := SQLStore{}
	service := UserService{store: db}
	
	service.WelcomeUser(101)
}
