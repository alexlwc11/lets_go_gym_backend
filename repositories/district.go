package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type DistrictRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) *DistrictRepository {
	return &DistrictRepository{DB: db}
}

func (dr *DistrictRepository) FindAll() ([]models.District, error) {
	var districts []models.District
	result := dr.DB.Find(&districts)

	if result.Error != nil {
		return nil, result.Error
	}

	return districts, nil
}
