package main

import (
	"os"

	dao "github.com/Wh0rigin/GraduationProject/dao"
	gin "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConifg()
	dao.InitDB()
	defer dao.CloseDb()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("servere.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConifg() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
