package repositories_test

import (
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"
	"time"

	"gorm.io/gorm"
)

// Mock repo with success result
type MockRegionRepositoryWithSuccessResult struct{}

func NewMockRegionRepositoryWithSuccessResult() repositories.RegionRepository {
	return &MockRegionRepositoryWithSuccessResult{}
}

func (m *MockRegionRepositoryWithSuccessResult) FindAll() (*[]models.Region, error) {
	result := []models.Region{}
	result = append(result, models.Region{
		BaseModel: models.BaseModel{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Code:   "T",
		NameEn: "TEST",
		NameZh: "TEST",
	})

	return &result, nil
}

// Mock repo with failure result
type MockRegionRepositoryWithFailureResult struct{}

func NewMockRegionRepositoryWithFailureResult() repositories.RegionRepository {
	return &MockRegionRepositoryWithFailureResult{}
}

func (m *MockRegionRepositoryWithFailureResult) FindAll() (*[]models.Region, error) {
	return &[]models.Region{}, gorm.ErrRecordNotFound
}
