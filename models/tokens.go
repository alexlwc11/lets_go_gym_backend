package models

import (
	"time"
)

type Token struct {
	BaseModel
	UserID    uint      `gorm:"<-:create;not null"`
	User      User      `gorm:"foreignKey:ID;references:UserID;constraint:OnUpdate:CASCADE" json:"-"`
	Value     string    `gorm:"<-:create;unique;not null"`
	ExpiredAt time.Time `gorm:"<-:create;not null"`
}

type SessionToken struct {
	Token
}

type RefreshToken struct {
	Token
}
