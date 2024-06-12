package apis

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppInfoHandler struct {
	AppVersionRepo repositories.AppVersionRepository
	DataInfoRepo   repositories.DataInfoRepository
}

func NewAppInfoHandler(
	appVersionRepo repositories.AppVersionRepository,
	dataInfoRepo repositories.DataInfoRepository,
) *AppInfoHandler {
	return &AppInfoHandler{AppVersionRepo: appVersionRepo, DataInfoRepo: dataInfoRepo}
}

func (aih *AppInfoHandler) RegisterRoutes(engine *gin.RouterGroup) {
	engine.GET("/app_info", aih.GetAppInfo)
}

// GetAppInfo godoc
//
//	@Summary		Get app info
//	@Description	Get latest app info
//	@Tags			AppInfo
//	@Produce		json
//	@Success		200	{object}	models.AppInfo
//	@Failure		500
//	@Router			/app_info [get]
func (aih *AppInfoHandler) GetAppInfo(c *gin.Context) {
	appVersion, err := aih.AppVersionRepo.FindAppVersion()
	if err != nil {
		log.Printf("Failed to find app version: %s\n", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	dataInfo, err := aih.DataInfoRepo.FindDataInfo()
	if err != nil {
		log.Printf("Failed to find data info: %s\n", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.AppInfo{
		AppVersion: *appVersion,
		DataInfo:   *dataInfo,
	})
}
