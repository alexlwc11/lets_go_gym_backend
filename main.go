package main

import (
	"lets_go_gym_backend/apis"
	"lets_go_gym_backend/configs"
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	mConfig, err := configs.InitConfig()
	if err != nil {
		panic("Failed to read config")
	}

	db, err := gorm.Open(mysql.Open(mConfig.GetDSNString()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Region{})
	db.AutoMigrate(&models.District{})
	db.AutoMigrate(&models.SportsCenter{})

	// region
	regionRepo := repositories.NewRegionRepository(db)
	regionHandler := apis.NewRegionHandler(regionRepo)

	// district
	districtRepo := repositories.NewDistrictRepository(db)
	districtHandler := apis.NewDistrictRepository(districtRepo)

	// sports center
	sportsCenterRepo := repositories.NewSportsCenterRepository(db)
	sportsCenterHandler := apis.NewSportsCenterRepository(sportsCenterRepo)

	router := gin.Default()
	// region
	router.GET("/regions", regionHandler.GetAllRegions)

	// district
	router.GET("/districts", districtHandler.GetAllDistricts)

	// sports center
	router.GET("/sports_centers", sportsCenterHandler.GetAllSportsCenters)

	router.Run("localhost:8080")
}
