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

func (rh *RegionHandler) RegisterRoutes(engine *gin.RouterGroup) {
	engine.GET("", rh.GetAllRegions)
}

type regionsOutDto struct {
	Region []models.Region `json:"regions"`
}

// GetAllRegions godoc
//
//	@Summary		Get all regions
//	@Description	Get latest regions data
//	@Tags			Regions
//	@Produce		json
//	@Success		200	{object}	regionsOutDto
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/regions [get]
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
