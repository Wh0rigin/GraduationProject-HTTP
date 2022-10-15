package po

import "gorm.io/gorm"

type SensorData struct {
	gorm.Model
	SensorType string
	Value      float32
}
