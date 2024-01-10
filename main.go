package main

import (
	"lets_go_gym_backend/configs"

	"gorm.io/gorm"
)

func main() {
	mConfig, err := configs.InitConfig()
	if err != nil {
		panic("Failed to read config")
	}

	db, err := gorm.Open(mConfig.GetDSNString(), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

}
