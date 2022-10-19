package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func SensorRouter(r *gin.Engine) {
	r.GET("/api/sensor/all/:type", middleware.AuthMiddleware(), controller.AllSensorDataController)
	r.GET("/api/sensor/recent/:type/:limit", middleware.AuthMiddleware(), controller.RecentSensorDataController)

	r.POST("/api/actuator/control", middleware.AuthMiddleware(), controller.ControlController)
}
