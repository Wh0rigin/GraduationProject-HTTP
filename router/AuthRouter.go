package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouther(r *gin.Engine) {
	//auth
	r.POST("/api/auth/register", controller.ResiterController)
	//use telephone
	r.POST("/api/auth/login", controller.LoginController)
	r.GET("/api/auth/detail", middleware.AuthMiddleware(), controller.DetailController)
}
