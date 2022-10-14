package dao

import (
	"fmt"

	"github.com/Wh0rigin/GraduationProject/bean"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDB *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "graduation"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	fmt.Printf(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//new table that is auto generate
	db.AutoMigrate(&bean.User{})
	myDB = db
}

func CloseDb() {
	sqldb, err := myDB.DB()
	if err != nil {
		panic("fail to close the Mysql")
	}
	sqldb.Close()
}

func GetDb() *gorm.DB {
	return myDB
}
