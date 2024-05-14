package repositories

import (
	models "lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateWithDeviceUUID(deviceUUID string) (models.User, error) {
	user := models.User{
		DeviceUUID: deviceUUID,
	}
	error := ur.DB.Create(&user).Error
	if error != nil {
		return models.User{}, error
	}

	return user, nil
}

func (ur *UserRepository) FindByDeviceUUID(deviceUUID string) (models.User, error) {
	var user models.User
	err := ur.DB.Model(models.User{DeviceUUID: deviceUUID}).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}