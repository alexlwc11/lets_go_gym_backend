package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type AppVersionRepository interface {
	FindAppVersion() (*models.AppVersion, error)
}

type AppVersionRepositoryImpl struct {
	DB *gorm.DB
}

func NewAppVersionRepositoryImpl(db *gorm.DB) AppVersionRepository {
	return &AppVersionRepositoryImpl{DB: db}
}

func (avr *AppVersionRepositoryImpl) FindAppVersion() (*models.AppVersion, error) {
	var appVersion models.AppVersion
	result := avr.DB.First(&appVersion)

	if result.Error != nil {
		return &models.AppVersion{}, result.Error
	}

	return &appVersion, nil
}
