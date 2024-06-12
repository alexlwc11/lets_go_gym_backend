package repositories_test

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
)

// Mock repo with success result
type MockSessionTokenRepository struct{}

func NewMockSessionTokenRepository() repositories.SessionTokenRepository {
	return &MockSessionTokenRepository{}
}

func (m *NewMockSessionTokenRepository) CreateWithUserId(userId uint) (*models.SessionToken, error) {

}
