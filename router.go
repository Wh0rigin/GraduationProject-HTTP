package main

import (
	"github.com/Wh0rigin/GraduationProject/router"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	router.AuthRouther(r)
	router.BookRouter(r)
	router.SensorRouter(r)
	return r
}
