package main

import (
	"io"
	"os"
	"time"

	"github.com/Wh0rigin/GraduationProject/common"
	_ "github.com/Wh0rigin/GraduationProject/docs"
	"github.com/Wh0rigin/GraduationProject/middleware"
	gin "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title Libary
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func main() {
	InitConifg()
	common.InitDB()

	// go func() {
	// 	for {
	// 		//TODO get newland
	// 	}
	// }()

	defer common.CloseDb()

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	// TODO session redis
	// store := common.Redisutil()
	// r.Use(sessions.Sessions("mysession", store))
	r = CollectRoute(r)
	port := viper.GetString("server.port")
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
	dateTime := time.Now().Format("2006_01_02_15_04_05")
	f, _ := os.Create("./log/startin" + dateTime + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
