package repositories_test

import (
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"

	"gorm.io/gorm"
)

// Mock repo with success result
type MockAppVersionRepositoryWithSuccessResult struct{}

func NewMockAppVersionRepositoryWithSuccessResult() repositories.AppVersionRepository {
	return &MockAppVersionRepositoryWithSuccessResult{}
}

func (m *MockAppVersionRepositoryWithSuccessResult) FindAppVersion() (*models.AppVersion, error) {
	return &models.AppVersion{
		LatestBuildVersion:  1,
		MinimumBuildVersion: 1,
	}, nil
}

// Mock repos with failure result
type MockAppVersionRepositoryWithFailureResult struct{}

func NewMockAppVersionRepositoryWithFailureResult() repositories.AppVersionRepository {
	return &MockAppVersionRepositoryWithFailureResult{}
}

func (m *MockAppVersionRepositoryWithFailureResult) FindAppVersion() (*models.AppVersion, error) {
	return &models.AppVersion{}, gorm.ErrRecordNotFound
}
