package dao

import (
	"github.com/Wh0rigin/GraduationProject/bean"
	"gorm.io/gorm"
)

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user bean.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
