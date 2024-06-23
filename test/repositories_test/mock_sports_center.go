package repositories_test

import (
	"fmt"
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"

	"gorm.io/gorm"
)

type MockSportsCenterRepositoryWithSuccessResult struct{}

func NewMockSportsCenterRepositoryWithSuccessResult() repositories.SportsCenterRepository {
	return &MockSportsCenterRepositoryWithSuccessResult{}
}

func (m *MockSportsCenterRepositoryWithSuccessResult) FindAll() (*[]models.SportsCenter, error) {
	hourlyQuota := int16(12)
	monthlyQuota := int16(24)
	result := []models.SportsCenter{}
	result = append(result, models.SportsCenter{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		ExternalID:   123,
		DistrictID:   1,
		NameEn:       "TESTING",
		NameZh:       "TESTING",
		AddressEn:    "TEST TEST TEST",
		AddressZh:    "TEST TEST TEST",
		PhoneNumbers: "12345678",
		HourlyQuota:  &hourlyQuota,
		MonthlyQuota: &monthlyQuota,
		LatitudeDMS:  "xx-xx-xx",
		LongitudeDMS: "xx-xx-xx",
		LatitudeDD:   "xxxxx",
		LongitudeDD:  "xxxxx",
	})

	return &result, nil
}

func (m *MockSportsCenterRepositoryWithSuccessResult) FindById(id uint) (*models.SportsCenter, error) {
	hourlyQuota := int16(12)
	monthlyQuota := int16(24)
	result := models.SportsCenter{
		BaseModel: models.BaseModel{
			ID: id,
		},
		ExternalID:   123,
		DistrictID:   1,
		NameEn:       "TESTING",
		NameZh:       "TESTING",
		AddressEn:    "TEST TEST TEST",
		AddressZh:    "TEST TEST TEST",
		PhoneNumbers: "12345678",
		HourlyQuota:  &hourlyQuota,
		MonthlyQuota: &monthlyQuota,
		LatitudeDMS:  "xx-xx-xx",
		LongitudeDMS: "xx-xx-xx",
		LatitudeDD:   "xxxxx",
		LongitudeDD:  "xxxxx",
	}

	return &result, nil
}

func (m *MockSportsCenterRepositoryWithSuccessResult) FindDetailsUrlById(id uint) (string, error) {
	sportsCenter, err := m.FindById(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=%d", sportsCenter.ExternalID), nil
}

type MockSportsCenterRepositoryWithFailureResult struct{}

func NewMockSportsCenterRepositoryWithFailureResult() repositories.SportsCenterRepository {
	return &MockSportsCenterRepositoryWithFailureResult{}
}

func (m *MockSportsCenterRepositoryWithFailureResult) FindAll() (*[]models.SportsCenter, error) {
	return &[]models.SportsCenter{}, gorm.ErrRecordNotFound
}

func (m *MockSportsCenterRepositoryWithFailureResult) FindById(id uint) (*models.SportsCenter, error) {
	return &models.SportsCenter{}, gorm.ErrRecordNotFound
}

func (m *MockSportsCenterRepositoryWithFailureResult) FindDetailsUrlById(id uint) (string, error) {
	sportsCenter, err := m.FindById(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=%d", sportsCenter.ExternalID), nil
}
