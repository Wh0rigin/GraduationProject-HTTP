package main

import (
	controller "github.com/Wh0rigin/GraduationProject/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.ResiterControler)
	return r
}
