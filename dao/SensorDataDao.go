package dao

import (
	"github.com/Wh0rigin/GraduationProject/po"
	"gorm.io/gorm"
)

func GetAllSensorDataByType(db *gorm.DB, stype string) (datas []po.SensorData) {
	db.Where("sensor_type = ?", stype).Find(&datas)
	return datas
}

func GetRecentSensorDataWithType(db *gorm.DB, stype string, limit int) (datas []po.SensorData) {
	db.Where("sensor_type = ?", stype).Order("id desc").Limit(limit).Find(&datas)
	return datas
}
