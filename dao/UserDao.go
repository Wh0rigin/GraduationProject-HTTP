package dao

import (
	"errors"

	"github.com/Wh0rigin/GraduationProject/bean"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user *bean.User) {
	db.Create(user)
}

func GetUser(db *gorm.DB, account string) (*bean.User, error) {
	var user bean.User
	db.Where("telephone=?", account).First(&user)
	if user.ID != 0 {
		return &user, nil
	}
	return nil, errors.New("dml:null value")
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user bean.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
