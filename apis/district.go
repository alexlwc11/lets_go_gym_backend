package apis

import (
	repositories "lets_go_gym_backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistrictHandler struct {
	DistrictRepo *repositories.DistrictRepository
}

func NewDistrictRepository(districtRepo *repositories.DistrictRepository) *DistrictHandler {
	return &DistrictHandler{DistrictRepo: districtRepo}
}

func (dh *DistrictHandler) GetAllDistricts(c *gin.Context) {
	districts, err := dh.DistrictRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, districts)
}
