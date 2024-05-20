package main

import (
	"lets_go_gym_backend/apis"
	"lets_go_gym_backend/configs"
	"lets_go_gym_backend/docs"
	LoadData "lets_go_gym_backend/load_data"
	"lets_go_gym_backend/middleware"
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "lets_go_gym_backend/docs"

	SwaggerFiles "github.com/swaggo/files"
	GinSwagger "github.com/swaggo/gin-swagger"
)

//	@title						Let's go gym API
//	@version					1.0
//	@description				Let's go gym API endpoints.
//
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@host						localhost:8080
//	@BasePath					/api/v1
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type in format of "Bearer --TOKEN--".

func main() {
	// init database connection
	db := initDatabase()

	// init router
	router := initRouter(db)

	router.Run("localhost:8080")
}

// Database
func initDatabase() *gorm.DB {
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
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.SessionToken{})
	db.AutoMigrate(&models.RefreshToken{})
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
			log.Println("Failed to load default regions data")
		}

		db.Create(defaultRegions)
	}

	if db.Take(&models.District{}).RowsAffected == 0 {
		defaultDistricts, err := LoadData.LoadDistricts("./data/districts.json")
		if err != nil {
			log.Println("Failed to load default districts data")
		}

		db.Create(defaultDistricts)
	}

	if db.Take(&models.SportsCenter{}).RowsAffected == 0 {
		defaultSportsCenters, err := LoadData.LoadSportsCenters("./data/sports_centers.json")
		if err != nil {
			log.Println("Failed to load default sports centers data")
		}

		db.Create(defaultSportsCenters)
	}
}

const (
	basePath = "/api/v1"
)

// Router
func initRouter(db *gorm.DB) *gin.Engine {
	// auth
	userRepo := repositories.NewUserRepository(db)
	sessionTokenRepo := repositories.NewSessionTokenRepository(db)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(db)
	authHandler := apis.NewAuthHandler(userRepo, sessionTokenRepo, refreshTokenRepo)

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
	router.SetTrustedProxies(nil)
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// auth
		setupAuthEndpoints(v1, authHandler)

		// app info
		setupAppInfoEndpoints(v1, appInfoHandler)

		// This middleware applies to the endpoints setup below
		v1.Use(middleware.AuthRequired(sessionTokenRepo.FindByValue))

		// region
		regions := v1.Group("/regions")
		setupRegionEndpoints(regions, regionHandler)

		// district
		districts := v1.Group("/districts")
		setupDistrictEndpoints(districts, districtHandler)

		// sports center
		sportsCenters := v1.Group("/sports_centers")
		setupSportsCenterEndpoints(sportsCenters, sportsCenterHandler)
	}

	router.GET("swagger/*any", GinSwagger.WrapHandler(SwaggerFiles.Handler))

	return router
}

// Setup endpoints
func setupAuthEndpoints(engine *gin.RouterGroup, authHandler *apis.AuthHandler) {
	engine.POST("/register", authHandler.Register)
	engine.POST("/sign_in", authHandler.SignIn)
	engine.POST("/refresh", authHandler.Refresh)
}

func setupAppInfoEndpoints(engine *gin.RouterGroup, appInfoHandler *apis.AppInfoHandler) {
	engine.GET("/app_info", appInfoHandler.GetAppInfo)
}

func setupRegionEndpoints(engine *gin.RouterGroup, regionHandler *apis.RegionHandler) {
	engine.GET("", regionHandler.GetAllRegions)
}

func setupDistrictEndpoints(engine *gin.RouterGroup, districtHandler *apis.DistrictHandler) {
	engine.GET("", districtHandler.GetAllDistricts)
}

func setupSportsCenterEndpoints(engine *gin.RouterGroup, sportsCenterHandler *apis.SportsCenterHandler) {
	engine.GET("", sportsCenterHandler.GetAllSportsCenters)
	engine.GET("/:id/details_url", sportsCenterHandler.GetDetailsUrl)
}
