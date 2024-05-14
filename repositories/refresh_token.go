package repositories

import (
	models "lets_go_gym_backend/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{DB: db}
}

// TODO move to env variable
const (
	refreshTokenValidTime = 90 * 24 * time.Hour
)

func (rtr *RefreshTokenRepository) CreateWithUserId(userId uint) (models.RefreshToken, error) {
	var expiredDuration = time.Now().Add(refreshTokenValidTime)

	hashedValue, err := bcrypt.GenerateFromPassword([]byte(uuid.NewString()), 4)
	if err != nil {
		return models.RefreshToken{}, err
	}

	tokenValue := string(hashedValue)
	token := models.RefreshToken{
		Token: models.Token{
			UserID:    userId,
			Value:     tokenValue,
			ExpiredAt: expiredDuration,
		},
	}

	error := rtr.DB.Create(&token).Error
	if error != nil {
		return models.RefreshToken{}, error
	}

	return token, nil
}

func (rtr *RefreshTokenRepository) FindByValue(value string) (models.RefreshToken, error) {
	var token models.RefreshToken
	err := rtr.DB.Where("value = ?", value).First(&token).Error
	if err != nil {
		return models.RefreshToken{}, err
	}

	return token, nil
}
