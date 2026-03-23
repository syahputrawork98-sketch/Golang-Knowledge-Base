package logic

import (
	"errors"
	"testing"
)

// 1. Definisikan Manual Mock
type MockRepo struct {
	getEmailFunc func(id int) (string, error)
}

func (m *MockRepo) GetUserEmail(id int) (string, error) {
	return m.getEmailFunc(id)
}

func TestGetWelcomeMessage(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		// 2. Inisialisasi Mock dengan perilaku sukses
		mock := &MockRepo{
			getEmailFunc: func(id int) (string, error) {
				return "test@example.com", nil
			},
		}
		svc := &Service{Repo: mock}

		msg, err := svc.GetWelcomeMessage(1)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if msg != "Welcome, user with email test@example.com!" {
			t.Errorf("unexpected message: %s", msg)
		}
	})

	t.Run("failure case from DB", func(t *testing.T) {
		// 3. Inisialisasi Mock dengan simulasi error
		mock := &MockRepo{
			getEmailFunc: func(id int) (string, error) {
				return "", errors.New("db connection timeout")
			},
		}
		svc := &Service{Repo: mock}

		_, err := svc.GetWelcomeMessage(1)
		if err == nil {
			t.Fatal("expected error but got none")
		}
	})
}
