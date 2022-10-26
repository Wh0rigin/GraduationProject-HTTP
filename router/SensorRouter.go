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
	r.GET("/api/sensor/all/non/filtrate", middleware.MangagerMiddleware(), controller.GetAllSensorDataNonFiltrate)
	//分页
	r.GET("/api/sensor/all/page", middleware.MangagerMiddleware(), controller.GetAllSensorDataLimit)

	r.GET("/api/sensor/total/:sensorName", middleware.MangagerMiddleware(), controller.SensorDataNumberByName)
}
