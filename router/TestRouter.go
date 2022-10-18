package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.Engine) {
	r.GET("/api/get/test", controller.TestGetContoller)
	r.POST("api/post/test200", controller.TestPostController200)
	r.POST("api/post/test422", controller.TestPostController422)
}
