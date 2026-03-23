package logic

import "fmt"

type Service struct {
	Repo Repository // Dependency Injection via Interface
}

func (s *Service) GetWelcomeMessage(id int) (string, error) {
	email, err := s.Repo.GetUserEmail(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Welcome, user with email %s!", email), nil
}
