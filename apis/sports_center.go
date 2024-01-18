package apis

import (
	"fmt"
	repositories "lets_go_gym_backend/repositories"
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

func (dh *SportsCenterHandler) GetAllSportsCenters(c *gin.Context) {
	sportsCenters, err := dh.SportsCenterRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"sports_centers": sportsCenters})
}

func (dh *SportsCenterHandler) GetDetailsUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	sportsCenter, err := dh.SportsCenterRepo.FindById(uint(id))
	print(sportsCenter.ExternalID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"url": generateDetailsUrl(sportsCenter.ExternalID)})
}

func generateDetailsUrl(externalId uint) string {
	return fmt.Sprintf("https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=%d", externalId)
}
