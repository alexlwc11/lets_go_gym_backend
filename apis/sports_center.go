package apis

import (
	repositories "lets_go_gym_backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SportsCenterHandler struct {
	SportsCenterRepo *repositories.SportsCenterRepository
}

func NewSportsCenterRepository(SportsCenterRepo *repositories.SportsCenterRepository) *SportsCenterHandler {
	return &SportsCenterHandler{SportsCenterRepo: SportsCenterRepo}
}

func (dh *SportsCenterHandler) GetAllSportsCenters(c *gin.Context) {
	sportsCenters, err := dh.SportsCenterRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, sportsCenters)
}
