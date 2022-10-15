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
	r.DELETE("/api/book/delete", middleware.AuthMiddleware(), controller.DeleteBookController)
	r.GET("/api/book/all", middleware.AuthMiddleware(), controller.AllBookController)
	//use ibsn
	r.POST("/api/book/add", middleware.AuthMiddleware(), controller.AddBookController)
	r.POST("/api/book/reduce", middleware.AuthMiddleware(), controller.ReduceBookController)
	r.POST("/api/book/rent", middleware.AuthMiddleware(), controller.RentBookController)
	r.POST("/api/book/return", middleware.AuthMiddleware(), controller.ReturnBookController)

	//api/book/select/isbn/{isbn}
	r.GET("/api/book/select/isbn/:isbn", middleware.AuthMiddleware(), controller.SelectBookController)
	return r
}
