package main

import "fmt"

// Antarmuka yang didefinisikan oleh Consumer
type UserRepository interface {
	Save(name string) error
}

// Logika Bisnis: Tidak peduli DB-nya apa
type UserService struct {
	repo UserRepository
}

func (s *UserService) Register(name string) {
	fmt.Printf("Registering user: %s\n", name)
	s.repo.Save(name)
}

// Implementasi Konkrit: MySQL
type MySQLRepo struct{}
func (m MySQLRepo) Save(name string) error {
	fmt.Println("Saved to MySQL Database")
	return nil
}

func main() {
	mysql := MySQLRepo{}
	
	// Dependency Injection: Masukkan MySQL ke Service
	service := UserService{repo: mysql}
	
	service.Register("Antigravity")
}
