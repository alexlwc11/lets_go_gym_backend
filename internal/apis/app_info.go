package apis

import (
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppInfoHandler interface {
	GetAppInfo(c *gin.Context)
}

type AppInfoHandlerImpl struct {
	AppVersionRepo repositories.AppVersionRepository
	DataInfoRepo   repositories.DataInfoRepository
}

func NewAppInfoHandlerImpl(
	appVersionRepo repositories.AppVersionRepository,
	dataInfoRepo repositories.DataInfoRepository,
) AppInfoHandler {
	return &AppInfoHandlerImpl{AppVersionRepo: appVersionRepo, DataInfoRepo: dataInfoRepo}
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
func (aihi *AppInfoHandlerImpl) GetAppInfo(c *gin.Context) {
	appVersion, err := aihi.AppVersionRepo.FindAppVersion()
	if err != nil {
		log.Printf("Failed to find app version: %s\n", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	dataInfo, err := aihi.DataInfoRepo.FindDataInfo()
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
