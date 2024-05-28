package main

import (
	"lets_go_gym_backend/apis"
	"lets_go_gym_backend/docs"
	"lets_go_gym_backend/middleware"
	"lets_go_gym_backend/repositories"

	"github.com/gin-gonic/gin"
	SwaggerFiles "github.com/swaggo/files"
	GinSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

const (
	basePath = "/api/v1"
)

type APIServer struct {
	addr   string
	db     *gorm.DB
	engine *gin.Engine
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Router
func (server *APIServer) Init() {
	// auth
	userRepo := repositories.NewUserRepositoryImpl(server.db)
	sessionTokenRepo := repositories.NewSessionTokenRepositoryImpl(server.db)
	refreshTokenRepo := repositories.NewRefreshTokenRepositoryImpl(server.db)
	authHandler := apis.NewAuthHandler(userRepo, sessionTokenRepo, refreshTokenRepo)

	// region
	regionRepo := repositories.NewRegionRepositoryImpl(server.db)
	regionHandler := apis.NewRegionHandler(regionRepo)

	// district
	districtRepo := repositories.NewDistrictRepositoryImpl(server.db)
	districtHandler := apis.NewDistrictHandler(districtRepo)

	// sports center
	sportsCenterRepo := repositories.NewSportsCenterRepositoryImpl(server.db)
	sportsCenterHandler := apis.NewSportsCenterHandler(sportsCenterRepo)

	// app info
	appVersionRepo := repositories.NewAppVersionRepositoryImpl(server.db)
	dataInfoRepo := repositories.NewDataInfoRepositoryImpl(server.db)
	appInfoHandler := apis.NewAppInfoHandler(appVersionRepo, dataInfoRepo)

	// setup router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	v1 := router.Group(basePath)
	{
		// auth
		authHandler.RegisterRoutes(v1)

		// app info
		appInfoHandler.RegisterRoutes(v1)

		// This middleware applies to the endpoints setup below
		v1.Use(middleware.AuthRequired(sessionTokenRepo.FindByValue))

		// region
		regions := v1.Group("/regions")
		regionHandler.RegisterRoutes(regions)

		// district
		districts := v1.Group("/districts")
		districtHandler.RegisterRoutes(districts)

		// sports center
		sportsCenters := v1.Group("/sports_centers")
		sportsCenterHandler.RegisterRoutes(sportsCenters)
	}

	server.engine = router
}

func (server *APIServer) SetupSwagger() {
	docs.SwaggerInfo.BasePath = basePath
	server.engine.GET("swagger/*any", GinSwagger.WrapHandler(SwaggerFiles.Handler))
}

func (server *APIServer) Run() {
	server.engine.Run(server.addr)
}
