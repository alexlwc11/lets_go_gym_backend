package repositories

import (
	"lets_go_gym_backend/internal/models"

	"gorm.io/gorm"
)

type RegionRepository interface {
	FindAll() (*[]models.Region, error)
}

type RegionRepositoryImpl struct {
	DB *gorm.DB
}

func NewRegionRepositoryImpl(db *gorm.DB) RegionRepository {
	return &RegionRepositoryImpl{DB: db}
}

func (rr *RegionRepositoryImpl) FindAll() (*[]models.Region, error) {
	var regions []models.Region
	result := rr.DB.Find(&regions)

	if result.Error != nil {
		return &[]models.Region{}, result.Error
	}

	return &regions, nil
}
