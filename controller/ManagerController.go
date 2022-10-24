package controller

import (
	"fmt"
	"net/http"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ManagerLogin godoc
// @Summary      Login an account
// @Description  login account by account and password return AccessToken
// @Tags         /api/manager
// @Accept       json
// @Produce      json
// @Param        account     body    string  true  "Account Telephone"
// @Param        password    body    string  true  "Account password"
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/manager/login [POST]
func ManagerLoginController(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	fmt.Println("{account:" + account + ",password:" + password + "}")

	if len(account) < 11 {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "电话长度必须为大于10位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "密码长度必须大于6位")
		return
	}

	user, err := dao.GetUserByAccount(dao.GetDb(), account)

	if user.Manager == false {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "您没有权限访问")
		return
	}

	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "账号错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "密码错误")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "Token生成失败")
		return
	}

	response.Response(ctx, http.StatusOK, 200,
		gin.H{
			"username":    user.Name,
			"AccessToken": token,
		}, "登录成功")
}
