package repositories

import (
	"lets_go_gym_backend/internal/models"

	"gorm.io/gorm"
)

type DistrictRepository interface {
	FindAll() (*[]models.District, error)
}

type DistrictRepositoryImpl struct {
	DB *gorm.DB
}

func NewDistrictRepositoryImpl(db *gorm.DB) DistrictRepository {
	return &DistrictRepositoryImpl{DB: db}
}

func (dr *DistrictRepositoryImpl) FindAll() (*[]models.District, error) {
	var districts []models.District
	result := dr.DB.Find(&districts)

	if result.Error != nil {
		return &[]models.District{}, result.Error
	}

	return &districts, nil
}
