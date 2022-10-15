package dto

import "github.com/Wh0rigin/GraduationProject/po"

type SensorDataDto struct {
	SensorName string  `Json:"SensorName"`
	Value      float32 `Json:"value"`
	RecordTime string  `Json:"recordTime"`
}

func NewSensorDataDto(data po.SensorData) *SensorDataDto {
	return &SensorDataDto{
		SensorName: data.SensorType,
		Value:      data.Value,
		RecordTime: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewSensorDataDtoByReference(data *po.SensorData) *SensorDataDto {
	return &SensorDataDto{
		SensorName: data.SensorType,
		Value:      data.Value,
		RecordTime: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewSensorDataDtoByArray(datas []po.SensorData) (sensorDataDtos []SensorDataDto) {
	sensorDataDtos = make([]SensorDataDto, len(datas))
	for i := 0; i < len(datas); i++ {
		sensorDataDtos[i] = *NewSensorDataDto(datas[i])
	}
	return sensorDataDtos
}
