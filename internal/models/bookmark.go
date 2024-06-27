package models

import "gorm.io/datatypes"

type UserBookmark struct {
	BaseModel
	UserID          uint           `gorm:"<-:create;uniqueIndex;not null"`
	SportsCenterIDs datatypes.JSON `gorm:"type:JSON"`
}
