package repositories

import (
	"lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type DataInfoRepository interface {
	FindDataInfo() (*models.DataInfo, error)
}

type DataInfoRepositoryImpl struct {
	DB *gorm.DB
}

func NewDataInfoRepositoryImpl(db *gorm.DB) DataInfoRepository {
	return &DataInfoRepositoryImpl{DB: db}
}

func (dir *DataInfoRepositoryImpl) FindDataInfo() (*models.DataInfo, error) {
	var dataInfo models.DataInfo
	result := dir.DB.Take(&dataInfo)
	if result.RowsAffected == 0 {
		newDataInfo, err := dir.createDataInfoRecord()
		if err != nil {
			return &models.DataInfo{}, err
		} else {
			dataInfo = *newDataInfo
		}
	}

	return &dataInfo, nil
}

func (dir *DataInfoRepositoryImpl) createDataInfoRecord() (*models.DataInfo, error) {
	var latestUpdatedRegion models.Region
	regionResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedRegion)
	if regionResult.Error != nil {
		println("Failed to get latest updated region")
		return &models.DataInfo{}, regionResult.Error
	}
	var latestUpdatedDistrict models.District
	districtResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedDistrict)
	if districtResult.Error != nil {
		println("Failed to get latest updated district")
		return &models.DataInfo{}, districtResult.Error
	}
	var latestUpdatedSportsCenter models.SportsCenter
	sportsCenterResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedSportsCenter)
	if sportsCenterResult.Error != nil {
		println("Failed to get latest updated sports center")
		return &models.DataInfo{}, sportsCenterResult.Error
	}

	var dataInfo = models.DataInfo{
		RegionDataLastUpdatedAt:       latestUpdatedRegion.UpdatedAt,
		DistrictDataLastUpdatedAt:     latestUpdatedDistrict.UpdatedAt,
		SportsCenterDataLastUpdatedAt: latestUpdatedSportsCenter.UpdatedAt,
	}

	dir.DB.Create(dataInfo)

	return &dataInfo, nil
}
