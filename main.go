package main

import (
	"os"

	dao "github.com/Wh0rigin/GraduationProject/dao"
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
	dao.InitDB()
	defer dao.CloseDb()
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
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
}
