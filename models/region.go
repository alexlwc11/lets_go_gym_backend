package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Code   string
	NameEn string
	NameZh string
}
