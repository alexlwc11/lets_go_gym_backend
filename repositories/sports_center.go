package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type SportsCenterRepository struct {
	DB *gorm.DB
}

func NewSportsCenterRepository(db *gorm.DB) *SportsCenterRepository {
	return &SportsCenterRepository{DB: db}
}

func (dr *SportsCenterRepository) FindAll() ([]models.SportsCenter, error) {
	var sportsCenters []models.SportsCenter
	result := dr.DB.Find(&sportsCenters)

	if result.Error != nil {
		return nil, result.Error
	}

	return sportsCenters, nil
}