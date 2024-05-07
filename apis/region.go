package apis

import (
	"lets_go_gym_backend/models"
	repositories "lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	RegionRepo *repositories.RegionRepository
}

func NewRegionHandler(regionRepo *repositories.RegionRepository) *RegionHandler {
	return &RegionHandler{RegionRepo: regionRepo}
}

type regionsOutDto struct {
	Region []models.Region `json:"regions"`
}

func (rh *RegionHandler) GetAllRegions(c *gin.Context) {
	regions, err := rh.RegionRepo.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, regionsOutDto{
		Region: regions,
	})
}
