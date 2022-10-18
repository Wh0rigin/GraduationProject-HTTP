package main

import (
	"net/http"

	"github.com/Wh0rigin/GraduationProject/response"
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
	port := viper.GetString("server.port")
	url := ginSwagger.URL("http://localhost:" + port + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/api/get/test", TestGetContoller)
	r.POST("api/post/test200", TestPostController200)
	r.POST("api/post/test422", TestPostController422)
	return r
}

func TestGetContoller(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 200, gin.H{"status": "Ok"}, "成功访问")
}

func TestPostController200(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 200, gin.H{"status": "Ok"}, "成功访问")
}

func TestPostController422(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 422, gin.H{"status": "error"}, "访问失败")
}
