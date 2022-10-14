package controller

import (
	"net/http"
	"time"

	"github.com/Wh0rigin/GraduationProject/bean"
	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginControler(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	if len(account) != 11 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "账号(电话号码)长度不正确",
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

	user, err := dao.GetUserByAccount(dao.GetDb(), account)
	if err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "账号错误",
			},
		)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422, "msg": "密码错误",
			},
		)
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": 500, "msg": "Token生成错误",
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "登录成功",
		"username":  user.Name,
		"ResultObj": gin.H{"AccessToken": token},
		"dateTime":  time.Now().String()[:19],
	})

}

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

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"code": 500, "msg": "密码加密出错",
			},
		)
		return
	}
	var user *bean.User = bean.NewUser(name, telephone, string(hashedpassword))

	dao.AddUser(dao.GetDb(), user)

	ctx.JSON(200, gin.H{
		"code":      200,
		"msg":       "注册成功",
		"name":      user.Name,
		"telephone": user.Telephone,
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK,
		gin.H{
			"code": 200,
			"data": gin.H{
				"user": dto.NewUserDto(user.(bean.User)),
			},
		})
}
