package main

import (
	"lets_go_gym_backend/apis"
	"lets_go_gym_backend/configs"
	LoadData "lets_go_gym_backend/load_data"
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// init database connection
	db := InitDatabase()

	// init router
	router := InitRouter(db)

	router.Run("localhost:8080")
}

// Database
func InitDatabase() *gorm.DB {
	mConfig, err := configs.InitConfig()
	if err != nil {
		panic("Failed to read config")
	}

	db, err := connectDatabase(mConfig)
	if err != nil {
		panic("Failed to connect database")
	}

	// Migration
	proceedSchemaMigration(db)

	// Load default data if needed
	loadDefaultDataIfNeeded(db)

	return db
}

func connectDatabase(config *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.GetDSNString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func proceedSchemaMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Region{})
	db.AutoMigrate(&models.District{})
	db.AutoMigrate(&models.SportsCenter{})
	db.AutoMigrate(&models.AppVersion{})
	db.AutoMigrate(&models.DataInfo{})
}

func loadDefaultDataIfNeeded(db *gorm.DB) {
	if db.Take(&models.Region{}).RowsAffected == 0 {
		defaultRegions, err := LoadData.LoadRegions("./data/regions.json")
		if err != nil {
			println("Failed to load default regions data")
		}

		db.Create(defaultRegions)
	}

	if db.Take(&models.District{}).RowsAffected == 0 {
		defaultDistricts, err := LoadData.LoadDistricts("./data/districts.json")
		if err != nil {
			println("Failed to load default districts data")
		}

		db.Create(defaultDistricts)
	}

	if db.Take(&models.SportsCenter{}).RowsAffected == 0 {
		defaultSportsCenters, err := LoadData.LoadSportsCenters("./data/sports_centers.json")
		if err != nil {
			println("Failed to load default sports centers data")
		}

		db.Create(defaultSportsCenters)
	}
}

// Router
func InitRouter(db *gorm.DB) *gin.Engine {
	// region
	regionRepo := repositories.NewRegionRepository(db)
	regionHandler := apis.NewRegionHandler(regionRepo)

	// district
	districtRepo := repositories.NewDistrictRepository(db)
	districtHandler := apis.NewDistrictRepository(districtRepo)

	// sports center
	sportsCenterRepo := repositories.NewSportsCenterRepository(db)
	sportsCenterHandler := apis.NewSportsCenterRepository(sportsCenterRepo)

	// app info
	appVersionRepo := repositories.NewAppVersionRepository(db)
	dataInfoRepo := repositories.NewDataInfoRepository(db)
	appInfoHandler := apis.NewAppInfoHandler(appVersionRepo, dataInfoRepo)

	router := gin.Default()
	// region
	setupRegionEndpoints(router, regionHandler)

	// district
	setupDistrictEndpoints(router, districtHandler)

	// sports center
	setupSportsCenterEndpoints(router, sportsCenterHandler)

	// app info
	setupAppInfoEndpoints(router, appInfoHandler)

	return router
}

func setupRegionEndpoints(engine *gin.Engine, regionHandler *apis.RegionHandler) {
	engine.GET("/regions", regionHandler.GetAllRegions)
}

func setupDistrictEndpoints(engine *gin.Engine, districtHandler *apis.DistrictHandler) {
	engine.GET("/districts", districtHandler.GetAllDistricts)
}

func setupSportsCenterEndpoints(engine *gin.Engine, sportsCenterHandler *apis.SportsCenterHandler) {
	engine.GET("/sports_centers", sportsCenterHandler.GetAllSportsCenters)
}

func setupAppInfoEndpoints(engine *gin.Engine, appInfoHandler *apis.AppInfoHandler) {
	engine.GET("/app_info", appInfoHandler.GetAppInfo)
}
