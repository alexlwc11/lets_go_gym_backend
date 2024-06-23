package repositories_test

import (
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"
	"time"

	"gorm.io/gorm"
)

// Mock repo with success result
type MockDistrictRepositoryWithSuccessResult struct{}

func NewMockDistrictRepositoryWithSuccessResult() repositories.DistrictRepository {
	return &MockDistrictRepositoryWithSuccessResult{}
}

func (m *MockDistrictRepositoryWithSuccessResult) FindAll() (*[]models.District, error) {
	result := []models.District{}
	result = append(result, models.District{
		BaseModel: models.BaseModel{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		RegionID: 1,
		NameEn:   "TEST",
		NameZh:   "TEST",
	})

	return &result, nil
}

// Mock repo with failure result
type MockDistrictRepositoryWithFailureResult struct{}

func NewMockDistrictRepositoryWithFailureResult() repositories.DistrictRepository {
	return &MockDistrictRepositoryWithFailureResult{}
}

func (m *MockDistrictRepositoryWithFailureResult) FindAll() (*[]models.District, error) {
	return &[]models.District{}, gorm.ErrRecordNotFound
}
