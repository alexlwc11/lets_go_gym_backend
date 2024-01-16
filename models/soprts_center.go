package models

import (
	"database/sql"
)

type SportsCenter struct {
	BaseModel
	ExternalID   uint          `gorm:"<-:create" json:"external_id"`
	DistrictID   uint          `json:"district_id"`
	District     District      `gorm:"<-:create;references:DistrictID"`
	NameEn       string        `json:"name_en"`
	NameZh       string        `json:"name_zh"`
	AddressEn    string        `json:"address_en"`
	AddressZh    string        `json:"address_zh"`
	PhoneNumbers []string      `json:"phone_numbers"`
	HourlyQuota  sql.NullInt16 `json:"hourly_quota"`
	MonthlyQuota sql.NullInt16 `json:"monthly_quota"`
	LatitudeDMS  string        `json:"latitude_dms"`
	LongitudeDMS string        `json:"longitude_dms"`
	LatitudeDD   string        `json:"latitude_dd"`
	LongitudeDD  string        `json:"longitude_dd"`
}
