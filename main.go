package main

import (
	"net/http"

	bean "github.com/Wh0rigin/GraduationProject/bean"
	dao "github.com/Wh0rigin/GraduationProject/dao"
	gin "github.com/gin-gonic/gin"
)

func main() {
	db := dao.InitDB()
	defer dao.CloseDb(db)
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		if len(telephone) != 11 {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"code": 422, "msg": "手机号必须为11位",
				},
			)
			return
		}
		if len(password) < 6 {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"code": 422, "msg": "密码不能少于6位",
				},
			)
			return
		}
		if dao.IsTelephoneExist(db, telephone) {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"code": 422, "msg": "用户手机号已经存在",
				},
			)
			return
		}

		var user *bean.User = bean.NewUser(name, telephone, password)

		dao.AddUser(db, user)

		ctx.JSON(200, gin.H{
			"msg":       "注册成功",
			"name":      user.Name,
			"telephone": user.Telephone,
		})
	})
	panic(r.Run())
}
