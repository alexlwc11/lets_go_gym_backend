package repositories

import (
	"encoding/json"
	"lets_go_gym_backend/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserBookmarkRepository interface {
	FindByUserId(userId uint) (*models.UserBookmark, error)
	UpdateWithUserId(userId uint, updatedSportsCenterIds []uint) error
}

type UserBookmarkRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserBookmarkRepositoryImpl(db *gorm.DB) UserBookmarkRepository {
	return &UserBookmarkRepositoryImpl{DB: db}
}

func (ubr *UserBookmarkRepositoryImpl) FindByUserId(userId uint) (*models.UserBookmark, error) {
	var userBookmark models.UserBookmark
	err := ubr.DB.Where("user_id = ?", userId).Attrs(models.UserBookmark{UserID: userId, SportsCenterIDs: datatypes.JSON("[]")}).FirstOrCreate(&userBookmark).Error
	if err != nil {
		return &models.UserBookmark{}, err
	}

	return &userBookmark, nil
}

func (ubr *UserBookmarkRepositoryImpl) UpdateWithUserId(userId uint, updatedSportsCenterIds []uint) error {
	updatedJSON, jsonErr := json.Marshal(updatedSportsCenterIds)
	if jsonErr != nil {
		return jsonErr
	}

	var userBookmark models.UserBookmark
	result := ubr.DB.Where("user_id = ?", userId).First(&userBookmark)
	if result.RowsAffected == 0 {
		userBookmark = models.UserBookmark{UserID: userId, SportsCenterIDs: datatypes.JSON(updatedJSON)}
	} else {
		userBookmark.SportsCenterIDs = datatypes.JSON(updatedJSON)
	}

	if err := ubr.DB.Save(&userBookmark).Error; err != nil {
		return err
	}

	return nil
}
