package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type SportsCenter struct {
	gorm.Model
	ExternalID   uint
	DistrictID   uint
	NameEn       string
	NameZh       string
	AddressEn    string
	AddressZh    string
	PhoneNumbers []string
	HourlyQuota  sql.NullInt16
	MonthlyQuota sql.NullInt16
	LatitudeDMS  string
	LongitudeDMS string
	LatitudeDD   string
	LongitudeDD  string
}
