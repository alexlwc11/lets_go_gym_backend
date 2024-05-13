package apis

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppInfoHandler struct {
	AppVersionRepo *repositories.AppVersionRepository
	DataInfoRepo   *repositories.DataInfoRepository
}

func NewAppInfoHandler(appVersionRepo *repositories.AppVersionRepository, dataInfoRepo *repositories.DataInfoRepository) *AppInfoHandler {
	return &AppInfoHandler{AppVersionRepo: appVersionRepo, DataInfoRepo: dataInfoRepo}
}

func (aih *AppInfoHandler) GetAppInfo(c *gin.Context) {
	appVersion, err := aih.AppVersionRepo.FindAppVersion()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	dataInfo, err := aih.DataInfoRepo.FindDataInfo()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.AppInfo{
		AppVersion: appVersion,
		DataInfo:   dataInfo,
	})
}
