package models

type Region struct {
	BaseModel
	Code   string `json:"code"`
	NameEn string `json:"name_en"`
	NameZh string `json:"name_zh"`
}
