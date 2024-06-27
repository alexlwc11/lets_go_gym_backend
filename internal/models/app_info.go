package models

import "time"

type AppInfo struct {
	AppVersion
	DataInfo DataInfo `json:"data_info"`
}

type AppVersion struct {
	LatestBuildVersion  uint `json:"latest_build_version"`
	MinimumBuildVersion uint `json:"minimum_build_version"`
}

func (AppVersion) TableName() string {
	return "app_version"
}

type DataInfo struct {
	RegionDataLastUpdatedAt       time.Time `json:"region_data_last_updated_at"`
	DistrictDataLastUpdatedAt     time.Time `json:"district_data_last_updated_at"`
	SportsCenterDataLastUpdatedAt time.Time `json:"sports_center_data_last_updated_at"`
}

func (DataInfo) TableName() string {
	return "data_info"
}
