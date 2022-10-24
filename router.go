package main

import (
	"github.com/Wh0rigin/GraduationProject/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	router.AuthRouther(r)
	router.BookRouter(r)
	router.SensorRouter(r)
	router.ManagerRouter(r)
	port := viper.GetString("server.port")
	url := ginSwagger.URL("http://localhost:" + port + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
