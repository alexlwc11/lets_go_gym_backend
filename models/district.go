package models

import "gorm.io/gorm"

type District struct {
	gorm.Model
	RegionID uint
	NameEn   string
	NameZh   string
}
