package main

import (
	controller "github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.ResiterController)
	r.POST("/api/auth/login", controller.LoginController)
	r.GET("/api/auth/detail", middleware.AuthMiddleware(), controller.DetailController)
	r.POST("/api/book/create", middleware.AuthMiddleware(), controller.CreateBookController)
	return r
}
