package repositories

import (
	"lets_go_gym_backend/models"

	"gorm.io/gorm"
)

type DataInfoRepository struct {
	DB *gorm.DB
}

func NewDataInfoRepository(db *gorm.DB) *DataInfoRepository {
	return &DataInfoRepository{DB: db}
}

func (dir *DataInfoRepository) FindDataInfo() (*models.DataInfo, error) {
	var dataInfo *models.DataInfo
	var err error
	result := dir.DB.Take(dataInfo)
	if result.RowsAffected == 0 {
		dataInfo, err = dir.createDataInfoRecord()
		if err != nil {
			return nil, err
		}
	}

	return dataInfo, nil
}

func (dir *DataInfoRepository) createDataInfoRecord() (*models.DataInfo, error) {
	var latestUpdatedRegion models.Region
	regionResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedRegion)
	if regionResult.Error != nil {
		println("Failed to get latest updated region")
		return nil, regionResult.Error
	}
	var latestUpdatedDistrict models.District
	districtResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedDistrict)
	if districtResult.Error != nil {
		println("Failed to get latest updated district")
		return nil, districtResult.Error
	}
	var latestUpdatedSportsCenter models.SportsCenter
	sportsCenterResult := dir.DB.Order("updated_at desc").Take(&latestUpdatedSportsCenter)
	if sportsCenterResult.Error != nil {
		println("Failed to get latest updated sports center")
		return nil, sportsCenterResult.Error
	}

	var dataInfo = models.DataInfo{
		RegionDataLastUpdatedAt:       latestUpdatedRegion.UpdatedAt,
		DistrictDataLastUpdatedAt:     latestUpdatedDistrict.UpdatedAt,
		SportsCenterDataLastUpdatedAt: latestUpdatedSportsCenter.UpdatedAt,
	}

	dir.DB.Create(dataInfo)

	return &dataInfo, nil
}
