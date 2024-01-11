package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type RegionRepository struct {
	DB *gorm.DB
}

func NewRegionRepository(db *gorm.DB) *RegionRepository {
	return &RegionRepository{DB: db}
}

func (rr *RegionRepository) FindAll() ([]models.Region, error) {
	var regions []models.Region
	result := rr.DB.Find(&regions)

	if result.Error != nil {
		return nil, result.Error
	}

	return regions, nil
}
