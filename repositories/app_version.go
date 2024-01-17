package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type AppVersionRepository struct {
	DB *gorm.DB
}

func NewAppVersionRepository(db *gorm.DB) *AppVersionRepository {
	return &AppVersionRepository{DB: db}
}

func (avr *AppVersionRepository) FindAppVersion() (*models.AppVersion, error) {
	var appVersion models.AppVersion
	result := avr.DB.First(&appVersion)

	if result.Error != nil {
		return nil, result.Error
	}

	return &appVersion, nil
}
