package repositories

import (
	"fmt"
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type SportsCenterRepository struct {
	DB *gorm.DB
}

func NewSportsCenterRepository(db *gorm.DB) *SportsCenterRepository {
	return &SportsCenterRepository{DB: db}
}

const (
	detailsUrl = "https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=%d"
)

func (dr *SportsCenterRepository) FindAll() ([]models.SportsCenter, error) {
	var sportsCenters []models.SportsCenter
	result := dr.DB.Find(&sportsCenters)

	if result.Error != nil {
		return []models.SportsCenter{}, result.Error
	}

	return sportsCenters, nil
}

func (dr *SportsCenterRepository) FindById(id uint) (models.SportsCenter, error) {
	var sportsCenter models.SportsCenter
	result := dr.DB.Where("id = ?", id).Take(&sportsCenter)

	if result.Error != nil {
		return models.SportsCenter{}, result.Error
	}

	return sportsCenter, nil
}

func (dr *SportsCenterRepository) FindDetailsUrlById(id uint) (string, error) {
	var sportsCenter models.SportsCenter
	result := dr.DB.Where("id = ?", id).Take(&sportsCenter)

	if result.Error != nil {
		return "", result.Error
	}

	return generateDetailsUrl(sportsCenter.ExternalID), nil
}

func generateDetailsUrl(externalId uint) string {
	return fmt.Sprintf(detailsUrl, externalId)
}
