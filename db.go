package main

import (
	"lets_go_gym_backend/configs"
	LoadData "lets_go_gym_backend/load_data"
	"lets_go_gym_backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	Auth "github.com/alexlwc11/simple_auth_go/cmd"
)

type SQLStorage struct {
	db *gorm.DB
}

func NewSQLStorage(config configs.Config) *SQLStorage {
	db, err := gorm.Open(mysql.Open(config.GetDSNString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL database!")

	return &SQLStorage{db: db}
}

func (storage *SQLStorage) Init() (*gorm.DB, error) {
	// Migration
	if err := proceedSchemaMigration(storage.db); err != nil {
		return nil, err
	}

	// Load default data if needed
	if err := loadDefaultDataIfNeeded(storage.db); err != nil {
		return nil, err
	}

	return storage.db, nil
}

func proceedSchemaMigration(db *gorm.DB) error {
	if err := Auth.ProceedSchemaMigration(db); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Region{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.District{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.SportsCenter{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.AppVersion{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.DataInfo{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.UserBookmark{}); err != nil {
		return err
	}

	return nil
}

func loadDefaultDataIfNeeded(db *gorm.DB) error {
	if db.Take(&models.Region{}).RowsAffected == 0 {
		defaultRegions, err := LoadData.LoadRegions("./data/regions.json")
		if err != nil {
			log.Println("Failed to load default regions data")
			return err
		}

		if err := db.Create(defaultRegions).Error; err != nil {
			log.Println("Failed to insert default regions data")
			return err
		}
	}

	if db.Take(&models.District{}).RowsAffected == 0 {
		defaultDistricts, err := LoadData.LoadDistricts("./data/districts.json")
		if err != nil {
			log.Println("Failed to load default districts data")
			return err
		}

		if err := db.Create(defaultDistricts).Error; err != nil {
			log.Println("Failed to insert default districts data")
			return err
		}
	}

	if db.Take(&models.SportsCenter{}).RowsAffected == 0 {
		defaultSportsCenters, err := LoadData.LoadSportsCenters("./data/sports_centers.json")
		if err != nil {
			log.Println("Failed to load default sports centers data")

			return err
		}

		if err := db.Create(defaultSportsCenters).Error; err != nil {
			log.Println("Failed to insert default sports centers data")
			return err
		}
	}

	return nil
}
