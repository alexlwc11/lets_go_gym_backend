package apis

import (
	"encoding/json"
	"lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserBookmarkHandler struct {
	UserBookmarkRepository repositories.UserBookmarkRepository
}

func NewUserBookmarkHandler(userBookmarkRepo repositories.UserBookmarkRepository) *UserBookmarkHandler {
	return &UserBookmarkHandler{UserBookmarkRepository: userBookmarkRepo}
}

func (ubh *UserBookmarkHandler) RegisterRoutes(engine *gin.RouterGroup) {
	engine.GET("", ubh.GetUserBookmarks)
	engine.PUT("", ubh.UpdateUserBookmarks)
}

// OutDto for [GetUserBookmarks]
type userBookmarkOutDto struct {
	SportsCenterIds []uint `json:"sports_center_ids"`
}

// GetUserBookmarks godoc
//
//	@Summary		GetUserBookmarks
//	@Description	Get user bookmarked sports centers
//	@Tags			Bookmarks
//	@Produce		json
//	@Success		200	{object}	userBookmarkOutDto
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/bookmarks [get]
func (ubh *UserBookmarkHandler) GetUserBookmarks(c *gin.Context) {
	userBookmark, err := ubh.UserBookmarkRepository.FindByUserId(c.GetUint("UserID"))
	// cannot find record for
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var sportsCenterIds []uint
	if jsonErr := json.Unmarshal(userBookmark.SportsCenterIDs, &sportsCenterIds); jsonErr != nil {
		log.Println(jsonErr.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, userBookmarkOutDto{SportsCenterIds: sportsCenterIds})
}

// InDto for [UpdateUserBookmarks]
type updateUserBookmarksInDto struct {
	UpdatedSportsCenterIds []uint `json:"updated_sports_center_ids"`
}

// UpdateUserBookmarks godoc
//
//	@Summary		UpdateUserBookmarks
//	@Description	Update user bookmarked sports centers
//	@Tags			Bookmarks
//	@Param			userUpdatedSportsCenterIds	body	updateUserBookmarksInDto	true	"Updated sports centers IDs"
//	@Success		200
//	@Failure		400
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/bookmarks [put]
func (ubh *UserBookmarkHandler) UpdateUserBookmarks(c *gin.Context) {
	var updateUserBookmarksInDto updateUserBookmarksInDto
	if err := c.BindJSON(&updateUserBookmarksInDto); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := ubh.UserBookmarkRepository.UpdateWithUserId(c.GetUint("UserID"), updateUserBookmarksInDto.UpdatedSportsCenterIds)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
