package models

type SportsCenter struct {
	BaseModel
	ExternalID   uint     `gorm:"<-:create" json:"external_id"`
	DistrictID   uint     `gorm:"<-:create" json:"district_id"`
	District     District `gorm:"foreignKey:ID;references:DistrictID;constraint:OnUpdate:CASCADE"`
	NameEn       string   `json:"name_en"`
	NameZh       string   `json:"name_zh"`
	AddressEn    string   `json:"address_en"`
	AddressZh    string   `json:"address_zh"`
	PhoneNumbers string   `json:"phone_numbers"`
	HourlyQuota  *int16   `json:"hourly_quota"`
	MonthlyQuota *int16   `json:"monthly_quota"`
	LatitudeDMS  string   `json:"latitude_dms"`
	LongitudeDMS string   `json:"longitude_dms"`
	LatitudeDD   string   `json:"latitude_dd"`
	LongitudeDD  string   `json:"longitude_dd"`
}
