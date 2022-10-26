package controller

import (
	"fmt"
	"net/http"
	"strconv"

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

	user, err := dao.GetUserByAccount(common.GetDb(), account)

	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "账号错误")
		return
	}

	if user.Manager == false {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "您没有权限访问")
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

	// TODO session Redis
	// session := sessions.Default(ctx)
	// session.Get("manager")

	// session.Set("manager", token)
	// err = session.Save()
	// if err != nil {
	// 	fmt.Println("Fail to save session:", err.Error())
	// } else {
	// 	fmt.Println("Succeed to save session")
	// }
	response.Response(ctx, http.StatusOK, 200,
		gin.H{
			"username":    user.Name,
			"AccessToken": token,
		}, "登录成功")
}

// logout godoc
// @Summary      logout
// @Description  login account by account and password return AccessToken
// @Tags         /api/auth
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/auth/logout [POST]
func LogoutController(ctx *gin.Context) {
	// TODO session Redis
	// session := sessions.Default(ctx)
	// // session.Delete("manager")
	// session.Set("manager", "") // this will mark the session as "written" and hopefully remove the username
	// // session.Set("manager", "") // this will mark the session as "written" and hopefully remove the username
	// session.Clear()
	// session.Options(sessions.Options{MaxAge: -1})
	// err := session.Save()
	// if err != nil {
	// 	fmt.Println("Fail to save session:", err.Error())
	// } else {
	// 	fmt.Println("Succeed to save session")
	// }
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "登出成功")
}

// GetAllUser godoc
// @Summary      GetAllUser
// @Description GetAllUser
// @Tags         /api/manager
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/manager/all/user [POST]
func GetAllUser(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "限制数量错误")
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "分页错误")
	}

	users := dao.GetAllUser(common.GetDb())
	if limit*page < len(users) {
		response.Response(ctx, http.StatusOK, 200, gin.H{
			"payload": users[limit*(page-1) : limit*page],
			"count":   len(users),
		}, "用户查询成功")
	} else {
		response.Response(ctx, http.StatusOK, 200, gin.H{
			"payload": users[limit*(page-1):],
			"count":   len(users),
		}, "用户查询成功")
	}

}

// DeleteUser godoc
// @Summary      DeleteUser
// @Description DeleteUser
// @Tags         /api/manager
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/manager/delete/user [POST]
func DeleteUser(ctx *gin.Context) {
	account := ctx.PostForm("account")
	user, err := dao.GetUserByAccount(common.GetDb(), account)
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "账号错误")
		return
	}
	dao.DeleteUser(common.GetDb(), user)
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "删除成功")
}

// UpdateUser godoc
// @Summary      UpdateUser
// @Description  UpdateUser
// @Tags         /api/manager
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJson
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/manager/update/user [POST]
func UpdateUser(ctx *gin.Context) {
	name := ctx.PostForm("user[Name]")
	manager := ctx.PostForm("user[Manager]")
	telephone := ctx.PostForm("user[Telephone]")
	password := ctx.PostForm("user[NewPassword]")

	for {
		userBean, err := dao.GetUserByAccount(common.GetDb(), name)

		if err != nil {
			response.Response(ctx, http.StatusOK, 202, gin.H{}, "没有这个用户")
			return
		}

		if manager == "true" {
			userBean.Manager = true
		} else {
			userBean.Manager = false
		}

		userBean.Name = name
		userBean.Telephone = telephone
		if password != "" {
			hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				response.Response(ctx, http.StatusOK, 500, gin.H{}, "密码加密失败")
				return
			}
			userBean.Password = string(hashedpassword)
		}

		uerr := dao.UpdateUser(common.GetDb(), userBean)
		if uerr == nil {
			break
		}
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "修改成功")
}
