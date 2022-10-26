package util

import "github.com/Wh0rigin/GraduationProject/po"

func ReverseSensorData(s []po.SensorData) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func AllNumber(s string) bool {
	for _, val := range s {
		if !(val >= '0' && val <= '9') {
			return false
		}
	}
	return true
}
