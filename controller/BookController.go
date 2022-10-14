package controller

import (
	"net/http"
	"strconv"

	"github.com/Wh0rigin/GraduationProject/bean"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
)

/*
*
param:

	name string
	isbn string
	number uint
	rentNumber uint
*/
func CreateBookController(ctx *gin.Context) {
	db := dao.GetDb()
	name := ctx.PostForm("name")
	isbn := ctx.PostForm("isbn")
	number, err := strconv.Atoi(ctx.PostForm("number"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "数量必须为数字")
		return
	}

	rentNumber, err := strconv.Atoi(ctx.PostForm("rentNumber"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "数量必须为数字")
		return
	}

	if name == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "名称不能为空")
		return
	}

	if rentNumber < 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "数量不能为负数")
		return
	}

	if number < 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "数量不能为负数")
		return
	}

	if dao.IsIsbnExist(db, isbn) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍已存在")
		return
	}

	book := bean.NewBookReference(name, isbn, uint(number), uint(rentNumber))
	dao.CreateBook(db, book)

	response.Response(ctx, http.StatusOK, 200, gin.H{}, "书本录入成功")
}
