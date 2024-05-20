package apis

import (
	"lets_go_gym_backend/models"
	repositories "lets_go_gym_backend/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SportsCenterHandler struct {
	SportsCenterRepo *repositories.SportsCenterRepository
}

func NewSportsCenterRepository(SportsCenterRepo *repositories.SportsCenterRepository) *SportsCenterHandler {
	return &SportsCenterHandler{SportsCenterRepo: SportsCenterRepo}
}

type sportsCentersOutDto struct {
	SportsCenters []models.SportsCenter `json:"sports_centers"`
}

func (dh *SportsCenterHandler) GetAllSportsCenters(c *gin.Context) {
	sportsCenters, err := dh.SportsCenterRepo.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, sportsCentersOutDto{
		SportsCenters: sportsCenters,
	})
}

type detailsUrlOutDto struct {
	Url string `json:"url"`
}

func (dh *SportsCenterHandler) GetDetailsUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	url, err := dh.SportsCenterRepo.FindDetailsUrlById(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, detailsUrlOutDto{
		Url: url,
	})
}
