package models

// TODO support other sign up methods e.g. email & password
type User struct {
	BaseModel
	DeviceUUID string `gorm:"uniqueIndex;size:36;not null"`
}
