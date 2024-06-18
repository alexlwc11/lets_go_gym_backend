package apis

import (
	"encoding/json"
	"lets_go_gym_backend/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserBookmarkHandler interface {
	GetUserBookmarks(c *gin.Context)
	PutUserBookmarks(c *gin.Context)
}

type UserBookmarkHandlerImpl struct {
	UserBookmarkRepository repositories.UserBookmarkRepository
}

func NewUserBookmarkHandlerImpl(userBookmarkRepo repositories.UserBookmarkRepository) UserBookmarkHandler {
	return &UserBookmarkHandlerImpl{UserBookmarkRepository: userBookmarkRepo}
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
func (ubh *UserBookmarkHandlerImpl) GetUserBookmarks(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Println("Cannot find user id")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userBookmark, err := ubh.UserBookmarkRepository.FindByUserId(userId.(uint))
	// cannot find record for the user
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var sportsCenterIds = []uint{}
	if userBookmark.SportsCenterIDs != nil {
		if jsonErr := json.Unmarshal(userBookmark.SportsCenterIDs, &sportsCenterIds); jsonErr != nil {
			log.Println(jsonErr.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, userBookmarkOutDto{SportsCenterIds: sportsCenterIds})
}

// InDto for [UpdateUserBookmarks]
type putUserBookmarksInDto struct {
	UpdatedSportsCenterIds []uint `json:"updated_sports_center_ids"`
}

// PutUserBookmarks godoc
//
//	@Summary		PutUserBookmarks
//	@Description	Update user bookmarked sports centers
//	@Tags			Bookmarks
//	@Param			userUpdatedSportsCenterIds	body	putUserBookmarksInDto	true	"Updated sports centers IDs"
//	@Success		200
//	@Failure		400
//	@Failure		403
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/bookmarks [put]
func (ubh *UserBookmarkHandlerImpl) PutUserBookmarks(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		log.Println("Cannot find user id")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var putUserBookmarksInDto putUserBookmarksInDto
	if err := c.BindJSON(&putUserBookmarksInDto); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := ubh.UserBookmarkRepository.UpdateWithUserId(userId.(uint), putUserBookmarksInDto.UpdatedSportsCenterIds)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
