package repositories_test

import (
	"encoding/json"
	"errors"
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
	"log"
	"time"

	"gorm.io/datatypes"
)

type MockUserBookmarkRepositoryWithSuccessResult struct {
	NextId   uint
	MockData []models.UserBookmark
}

func NewMockUserBookmarkRepositoryWithSuccessResult() repositories.UserBookmarkRepository {
	createdAt := time.Now()
	emptyIds := []uint{}
	ids := []uint{1, 2, 3}
	emptyIdsJson, _ := json.Marshal(emptyIds)
	idsJson, _ := json.Marshal(ids)
	return &MockUserBookmarkRepositoryWithSuccessResult{
		NextId: 3,
		MockData: []models.UserBookmark{
			{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
				UserID:          1,
				SportsCenterIDs: datatypes.JSON(emptyIdsJson),
			},
			{
				BaseModel: models.BaseModel{
					ID:        2,
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				},
				UserID:          4,
				SportsCenterIDs: datatypes.JSON(idsJson),
			},
		},
	}
}

func (m *MockUserBookmarkRepositoryWithSuccessResult) FindByUserId(userId uint) (*models.UserBookmark, error) {
	for _, data := range m.MockData {
		if data.UserID == userId {
			return &data, nil
		}
	}

	return &models.UserBookmark{}, nil
}

func (m *MockUserBookmarkRepositoryWithSuccessResult) UpdateWithUserId(userId uint, updatedSportsCenterIds []uint) error {
	idsJson, _ := json.Marshal(updatedSportsCenterIds)
	for index, data := range m.MockData {
		if data.UserID == userId {
			// update existing record
			log.Println("Update existing record")
			m.MockData[index].SportsCenterIDs = datatypes.JSON(idsJson)
			m.MockData[index].UpdatedAt = time.Now()
			return nil
		}
	}

	// add new record
	log.Println("Create new record")
	createdAt := time.Now()
	m.MockData = append(m.MockData, models.UserBookmark{
		BaseModel: models.BaseModel{
			ID:        m.NextId,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		},
		UserID:          userId,
		SportsCenterIDs: datatypes.JSON(idsJson),
	})
	m.NextId = m.NextId + 1

	return nil
}

type MockUserBookmarkRepositoryWithFailureResult struct{}

func NewMockUserBookmarkRepositoryWithFailureResult() repositories.UserBookmarkRepository {
	return &MockUserBookmarkRepositoryWithFailureResult{}
}

func (m *MockUserBookmarkRepositoryWithFailureResult) FindByUserId(userId uint) (*models.UserBookmark, error) {
	return &models.UserBookmark{}, errors.New("Mocking failure result")
}

func (m *MockUserBookmarkRepositoryWithFailureResult) UpdateWithUserId(userId uint, updatedSportsCenterIds []uint) error {
	return errors.New("Mocking failure result")
}
