package repositories

import (
	models "lets_go_gym_backend/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SessionTokenRepository interface {
	CreateWithUserId(userId uint) (*models.SessionToken, error)
	FindByValue(value string) (*models.SessionToken, error)
}

type SessionTokenRepositoryImpl struct {
	DB *gorm.DB
}

func NewSessionTokenRepositoryImpl(db *gorm.DB) SessionTokenRepository {
	return &SessionTokenRepositoryImpl{DB: db}
}

// TODO move to env variable
const (
	sessionTokenValidTime = 2 * 24 * time.Hour
)

func (str *SessionTokenRepositoryImpl) CreateWithUserId(userId uint) (*models.SessionToken, error) {
	var expiredDuration = time.Now().Add(sessionTokenValidTime)

	hashedValue, err := bcrypt.GenerateFromPassword([]byte(uuid.NewString()), 4)
	if err != nil {
		return &models.SessionToken{}, err
	}

	tokenValue := string(hashedValue)
	token := models.SessionToken{
		Token: models.Token{
			UserID:    userId,
			Value:     tokenValue,
			ExpiredAt: expiredDuration,
		},
	}

	dbError := str.DB.Create(&token).Error
	if dbError != nil {
		return &models.SessionToken{}, dbError
	}

	return &token, nil
}

func (str *SessionTokenRepositoryImpl) FindByValue(value string) (*models.SessionToken, error) {
	var token models.SessionToken
	err := str.DB.Where("value = ?", value).First(&token).Error
	if err != nil {
		return &models.SessionToken{}, err
	}

	return &token, nil
}
