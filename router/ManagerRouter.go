package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/gin-gonic/gin"
)

func ManagerRouter(r *gin.Engine) {
	r.POST("/api/manager/login", controller.ManagerLoginController)
	r.POST("/api/manager/logout", controller.LogoutController)
}
