package apis

import (
	"lets_go_gym_backend/internal/models"
	"lets_go_gym_backend/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistrictHandler interface {
	GetAllDistricts(c *gin.Context)
}

type DistrictHandlerImpl struct {
	DistrictRepo repositories.DistrictRepository
}

func NewDistrictHandlerImpl(districtRepo repositories.DistrictRepository) DistrictHandler {
	return &DistrictHandlerImpl{DistrictRepo: districtRepo}
}

type districtsOutDto struct {
	Districts []models.District `json:"districts"`
}

// GetAllDistricts godoc
//
//	@Summary		Get all districts
//	@Description	Get latest districts data
//	@Tags			Districts
//	@Produce		json
//	@Success		200	{object}	districtsOutDto
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/districts [get]
func (dh *DistrictHandlerImpl) GetAllDistricts(c *gin.Context) {
	districts, err := dh.DistrictRepo.FindAll()
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, districtsOutDto{
		Districts: *districts,
	})
}
