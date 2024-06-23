package api

import (
	"lets_go_gym_backend/internal/apis"
	"lets_go_gym_backend/internal/docs"
	"lets_go_gym_backend/internal/repositories"

	"github.com/gin-gonic/gin"
	SwaggerFiles "github.com/swaggo/files"
	GinSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	Auth "github.com/alexlwc11/simple_auth_go/cmd"
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
	authHandler := Auth.CreateAuthHandler(server.db)

	// app info
	appInfoHandler := apis.NewAppInfoHandlerImpl(
		repositories.NewAppVersionRepositoryImpl(server.db),
		repositories.NewDataInfoRepositoryImpl(server.db),
	)

	// region
	regionHandler := apis.NewRegionHandlerImpl(
		repositories.NewRegionRepositoryImpl(server.db),
	)

	// district
	districtHandler := apis.NewDistrictHandlerImpl(
		repositories.NewDistrictRepositoryImpl(server.db),
	)

	// sports center
	sportsCenterHandler := apis.NewSportsCenterHandlerImpl(
		repositories.NewSportsCenterRepositoryImpl(server.db),
	)

	// bookmark
	userBookmarkHandler := apis.NewUserBookmarkHandlerImpl(
		repositories.NewUserBookmarkRepositoryImpl(server.db),
	)

	// setup router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	v1 := router.Group(basePath)
	{
		// auth
		v1.POST("/register", authHandler.Register)
		v1.POST("/sign_in", authHandler.SignIn)
		v1.POST("/refresh", authHandler.Refresh)

		// app info
		v1.GET("/app_info", appInfoHandler.GetAppInfo)

		// This middleware applies to the endpoints setup below
		v1.Use(Auth.AuthRequired(server.db))

		// region
		regions := v1.Group("/regions")
		regions.GET("", regionHandler.GetAllRegions)

		// district
		districts := v1.Group("/districts")
		districts.GET("", districtHandler.GetAllDistricts)

		// sports center
		sportsCenters := v1.Group("/sports_centers")
		sportsCenters.GET("", sportsCenterHandler.GetAllSportsCenters)
		sportsCenters.GET("/:id/details_url", sportsCenterHandler.GetDetailsUrl)

		// bookmark
		bookmarks := v1.Group("/bookmarks")
		bookmarks.GET("", userBookmarkHandler.GetUserBookmarks)
		bookmarks.PUT("", userBookmarkHandler.PutUserBookmarks)
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