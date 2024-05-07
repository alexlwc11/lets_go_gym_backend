package apis

import (
	"lets_go_gym_backend/models"
	repositories "lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistrictHandler struct {
	DistrictRepo *repositories.DistrictRepository
}

func NewDistrictRepository(districtRepo *repositories.DistrictRepository) *DistrictHandler {
	return &DistrictHandler{DistrictRepo: districtRepo}
}

type districtsOutDto struct {
	Districts []models.District `json:"districts"`
}

func (dh *DistrictHandler) GetAllDistricts(c *gin.Context) {
	districts, err := dh.DistrictRepo.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, districtsOutDto{
		Districts: districts,
	})
}
