package dao

import (
	"errors"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/po"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user *po.User) {
	db.Create(user)
}

func GetUserByAccount(db *gorm.DB, account string) (*po.User, error) {
	var user po.User
	db.Where("telephone=?", account).First(&user)
	if user.ID != 0 {
		return &user, nil
	}
	return nil, errors.New("dml:null value")
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user po.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func GetUserById(db *gorm.DB, user po.User, uid uint) po.User {
	db.First(&user, uid)
	return user
}

func GetAllUser(db *gorm.DB) (users []po.User) {
	db.Find(&users)
	return users
}

func DeleteUser(db *gorm.DB, user *po.User) {
	db.Delete(user)
}

func UpdateUser(db *gorm.DB, user *po.User) error {
	var nowUser po.User
	nowUser = GetUserById(common.GetDb(), nowUser, user.ID)

	// positive lock
	if nowUser.Version != user.Version {
		return errors.New("dml:version clash")
	}

	user.Version += 1
	db.Save(user)
	return nil
}
