package controller

import (
	"net/http"

	"github.com/Wh0rigin/GraduationProject/bean"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/gin-gonic/gin"
)

func ResiterControler(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	if len(telephone) != 11 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "电话长度不对",
			},
		)
		return
	}
	if len(password) < 6 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "密码长度不够",
			},
		)
		return
	}
	if dao.IsTelephoneExist(dao.GetDb(), telephone) {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "电话已存在",
			},
		)
		return
	}

	var user *bean.User = bean.NewUser(name, telephone, password)

	dao.AddUser(dao.GetDb(), user)

	ctx.JSON(200, gin.H{
		"msg":       "注册成功",
		"name":      user.Name,
		"telephone": user.Telephone,
	})
}
