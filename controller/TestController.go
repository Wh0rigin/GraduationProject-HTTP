package controller

import (
	"net/http"

	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
)

func TestGetContoller(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 200, gin.H{"status": "Ok"}, "成功访问")
}

func TestPostController200(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 200, gin.H{"status": "Ok"}, "成功访问")
}

func TestPostController422(ctx *gin.Context) {
	response.Response(ctx, http.StatusOK, 422, gin.H{"status": "error"}, "访问失败")
}
