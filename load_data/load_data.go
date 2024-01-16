package load_data

import (
	"encoding/json"
	"io"
	"lets_go_gym_backend/models"
	"os"
)

type Regions struct {
	Regions []models.Region `json:"regions"`
}

type Districts struct {
	Districts []models.District `json:"districts"`
}

type SportsCenters struct {
	SportsCenters []models.SportsCenter `json:"sports_centers"`
}

func LoadRegions(filePath string) ([]models.Region, error) {
	regions, err := loadDataFromJson[Regions](filePath)
	if err != nil {
		return nil, err
	}
	return regions.Regions, nil
}

func LoadDistricts(filePath string) ([]models.District, error) {
	districts, err := loadDataFromJson[Districts](filePath)
	if err != nil {
		return nil, err
	}
	return districts.Districts, nil
}

func LoadSportsCenters(filePath string) ([]models.SportsCenter, error) {
	sportsCenters, err := loadDataFromJson[SportsCenters](filePath)
	if err != nil {
		return nil, err
	}
	return sportsCenters.SportsCenters, nil
}

func loadDataFromJson[V Regions | Districts | SportsCenters](path string) (*V, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		println(err.Error())
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result V
	json.Unmarshal([]byte(byteValue), &result)

	return &result, nil
}
