package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func ManagerRouter(r *gin.Engine) {
	r.POST("/api/manager/login", controller.ManagerLoginController)
	r.POST("/api/manager/logout", controller.LogoutController)
	r.GET("/api/manager/all/user", middleware.MangagerMiddleware(), controller.GetAllUser)
	r.POST("/api/manager/delete/user", middleware.MangagerMiddleware(), controller.DeleteUser)
	r.POST("/api/manager/update/user", middleware.MangagerMiddleware(), controller.UpdateUser)
}
