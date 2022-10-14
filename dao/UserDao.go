package dao

import (
	"github.com/Wh0rigin/GraduationProject/bean"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user *bean.User) {
	db.Create(user)
}
