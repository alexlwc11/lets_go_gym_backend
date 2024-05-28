package apis

import (
	"lets_go_gym_backend/models"
	repositories "lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistrictHandler struct {
	DistrictRepo repositories.DistrictRepository
}

func NewDistrictHandler(districtRepo repositories.DistrictRepository) *DistrictHandler {
	return &DistrictHandler{DistrictRepo: districtRepo}
}

func (dh *DistrictHandler) RegisterRoutes(engine *gin.RouterGroup) {
	engine.GET("", dh.GetAllDistricts)
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
func (dh *DistrictHandler) GetAllDistricts(c *gin.Context) {
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
