package main

import (
	controller "github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.ResiterControler)
	r.POST("/api/auth/login", controller.LoginControler)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
