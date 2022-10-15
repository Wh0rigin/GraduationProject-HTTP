package main

import (
	controller "github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//auth
	r.POST("/api/auth/register", controller.ResiterController)
	//use telephone
	r.POST("/api/auth/login", controller.LoginController)
	r.GET("/api/auth/detail", middleware.AuthMiddleware(), controller.DetailController)
	//book
	r.POST("/api/book/create", middleware.AuthMiddleware(), controller.CreateBookController)
	//user ibsn
	r.POST("/api/book/add", middleware.AuthMiddleware(), controller.AddBookController)
	r.POST("api/book/reduce", middleware.AuthMiddleware(), controller.ReduceBookController)
	return r
}
