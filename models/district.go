package models

type District struct {
	BaseModel
	RegionID uint   `gorm:"<-:create" json:"region_id"`
	Region   Region `gorm:"foreignKey:ID;references:RegionID;constraint:OnUpdate:CASCADE"`
	NameEn   string `json:"name_en"`
	NameZh   string `json:"name_zh"`
}
