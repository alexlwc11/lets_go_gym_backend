package apis

import (
	"lets_go_gym_backend/models"
	"lets_go_gym_backend/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SportsCenterHandler interface {
	GetAllSportsCenters(c *gin.Context)
	GetDetailsUrl(c *gin.Context)
}

type SportsCenterHandlerImpl struct {
	SportsCenterRepo repositories.SportsCenterRepository
}

func NewSportsCenterHandlerImpl(SportsCenterRepo repositories.SportsCenterRepository) SportsCenterHandler {
	return &SportsCenterHandlerImpl{SportsCenterRepo: SportsCenterRepo}
}

type sportsCentersOutDto struct {
	SportsCenters []models.SportsCenter `json:"sports_centers"`
}

// GetAllSportsCenters godoc
//
//	@Summary		Get all sports centers
//	@Description	getting latest sports centers data
//	@Tags			Sports centers
//	@Produce		json
//	@Success		200	{object}	sportsCentersOutDto
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/sports_centers [get]
func (sch *SportsCenterHandlerImpl) GetAllSportsCenters(c *gin.Context) {
	sportsCenters, err := sch.SportsCenterRepo.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, sportsCentersOutDto{
		SportsCenters: *sportsCenters,
	})
}

type detailsUrlOutDto struct {
	Url string `json:"url"`
}

// GetDetailsUrl godoc
//
//	@Summary		Get details url
//	@Description	Get the details url for specified sports center
//	@Tags			Sports centers
//	@Produce		json
//	@Param			id	path		string	true	"Sports center ID"
//	@Success		200	{object}	detailsUrlOutDto
//	@Failure		404
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/sports_centers/{id}/details_url [get]
func (sch *SportsCenterHandlerImpl) GetDetailsUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	url, err := sch.SportsCenterRepo.FindDetailsUrlById(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, detailsUrlOutDto{
		Url: url,
	})
}
