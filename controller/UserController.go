package controller

import (
	"net/http"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/Wh0rigin/GraduationProject/po"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary      Login an account
// @Description  login account by account and password return AccessToken
// @Tags         /api/auth
// @Accept       json
// @Produce      json
// @Param        account     body    string  true  "Account Telephone"
// @Param        password    body    string  true  "Account password"
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/auth/login [POST]
func LoginController(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	if len(account) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "电话长度必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码长度必须大于6位")
		return
	}

	user, err := dao.GetUserByAccount(dao.GetDb(), account)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "账号错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码错误")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "Token生成失败")
		return
	}
	response.Response(ctx, http.StatusOK, 200,
		gin.H{
			"username":    user.Name,
			"AccessToken": token,
		}, "登录成功")
}

// Resiter godoc
// @Summary      Resiter an account
// @Description  create account by account and password
// @Tags         /api/auth
// @Accept       json
// @Produce      json
// @Param        name     body    string  true  "Account Name"
// @Param        account     body    string  true  "Account Telephone"
// @Param        password    body    string  true  "Account password"
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/auth/resiter [POST]
func ResiterController(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "电话长度必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码长度必须大于6位")
		return
	}
	if dao.IsTelephoneExist(dao.GetDb(), telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "电话号码已存在")
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "密码加密失败")
		return
	}
	var user po.User = po.NewUser(name, telephone, string(hashedpassword))

	dao.AddUser(dao.GetDb(), &user)

	response.Response(ctx, http.StatusOK, 200, gin.H{"user": dto.NewUserDto(user)}, "注册成功")
}

// Resiter godoc
// @Summary      Resiter an account
// @Description  create account by account and password
// @Tags         /api/auth
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/auth/datail [GET]
func DetailController(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Response(ctx, http.StatusOK, 200, gin.H{"user": dto.NewUserDto(user.(po.User))}, "查询成功")
}
