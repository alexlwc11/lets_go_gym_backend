package repositories

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/utils"
	"time"

	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	CreateWithUserId(userId uint) (*models.RefreshToken, error)
	FindByValue(value string) (*models.RefreshToken, error)
}

type RefreshTokenRepositoryImpl struct {
	DB *gorm.DB
}

func NewRefreshTokenRepositoryImpl(db *gorm.DB) RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{DB: db}
}

// TODO move to env variable
const (
	refreshTokenValidTime = 90 * 24 * time.Hour
)

func (rtr *RefreshTokenRepositoryImpl) CreateWithUserId(userId uint) (*models.RefreshToken, error) {
	expiredDuration := time.Now().Add(refreshTokenValidTime)

	tokenValue, err := utils.GenerateToken()
	if err != nil {
		return &models.RefreshToken{}, err
	}

	token := models.RefreshToken{
		Token: models.Token{
			UserID:    userId,
			Value:     tokenValue,
			ExpiredAt: expiredDuration,
		},
	}

	error := rtr.DB.Create(&token).Error
	if error != nil {
		return &models.RefreshToken{}, error
	}

	return &token, nil
}

func (rtr *RefreshTokenRepositoryImpl) FindByValue(value string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := rtr.DB.Where("value = ?", value).First(&token).Error
	if err != nil {
		return &models.RefreshToken{}, err
	}

	return &token, nil
}
