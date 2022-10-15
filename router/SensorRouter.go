package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/gin-gonic/gin"
)

func SensorRouter(r *gin.Engine) {
	r.GET("/api/sensor/all/:type", controller.AllSensorDataController)
	r.GET("/api/sensor/recent/:type/:limit", controller.RecentSensorDataController)
}
