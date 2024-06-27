package repositories

import (
	"fmt"
	"lets_go_gym_backend/internal/models"

	"gorm.io/gorm"
)

type SportsCenterRepository interface {
	FindAll() (*[]models.SportsCenter, error)
	FindById(id uint) (*models.SportsCenter, error)
	FindDetailsUrlById(id uint) (string, error)
}

type SportsCenterRepositoryImpl struct {
	DB *gorm.DB
}

func NewSportsCenterRepositoryImpl(db *gorm.DB) SportsCenterRepository {
	return &SportsCenterRepositoryImpl{DB: db}
}

const (
	detailsUrl = "https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=%d"
)

func (dr *SportsCenterRepositoryImpl) FindAll() (*[]models.SportsCenter, error) {
	var sportsCenters []models.SportsCenter
	result := dr.DB.Find(&sportsCenters)

	if result.Error != nil {
		return &[]models.SportsCenter{}, result.Error
	}

	return &sportsCenters, nil
}

func (dr *SportsCenterRepositoryImpl) FindById(id uint) (*models.SportsCenter, error) {
	var sportsCenter models.SportsCenter
	result := dr.DB.Where("id = ?", id).Take(&sportsCenter)

	if result.Error != nil {
		return &models.SportsCenter{}, result.Error
	}

	return &sportsCenter, nil
}

func (dr *SportsCenterRepositoryImpl) FindDetailsUrlById(id uint) (string, error) {
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
