package apis

import (
	repositories "lets_go_gym_backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	RegionRepo *repositories.RegionRepository
}

func NewRegionHandler(regionRepo *repositories.RegionRepository) *RegionHandler {
	return &RegionHandler{RegionRepo: regionRepo}
}

func (rh *RegionHandler) GetAllRegions(c *gin.Context) {
	regions, err := rh.RegionRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regions)
}
