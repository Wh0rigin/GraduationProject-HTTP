package dao

import (
	"fmt"

	"github.com/Wh0rigin/GraduationProject/bean"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDB *gorm.DB

func InitDB() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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
	db.AutoMigrate(&bean.Book{})
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
