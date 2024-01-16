package models

type District struct {
	BaseModel
	RegionID uint   `json:"region_id"`
	Region   Region `gorm:"<-:create;references:RegionID"`
	NameEn   string `json:"name_en"`
	NameZh   string `json:"name_zh"`
}
