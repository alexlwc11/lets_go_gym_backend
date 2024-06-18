package repositories_test

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
	"time"

	"gorm.io/gorm"
)

// Mock repo with success result
type MockDataInfoRepositoryWithSuccessResult struct{}

func NewMockDataInfoRepositoryWithSuccessResult() repositories.DataInfoRepository {
	return &MockDataInfoRepositoryWithSuccessResult{}
}

func (m *MockDataInfoRepositoryWithSuccessResult) FindDataInfo() (*models.DataInfo, error) {
	return &models.DataInfo{
		RegionDataLastUpdatedAt:       time.Now(),
		DistrictDataLastUpdatedAt:     time.Now(),
		SportsCenterDataLastUpdatedAt: time.Now(),
	}, nil
}

// Mock repo with failure result
type MockDataInfoRepositoryWithFailureResult struct{}

func NewMockDataInfoRepositoryWithFailureResult() repositories.DataInfoRepository {
	return &MockDataInfoRepositoryWithFailureResult{}
}

func (m *MockDataInfoRepositoryWithFailureResult) FindDataInfo() (*models.DataInfo, error) {
	return &models.DataInfo{}, gorm.ErrRecordNotFound
}
