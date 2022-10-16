package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/Wh0rigin/GraduationProject/po"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
)

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

	if rentNumber < 0 || number < 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "数量不能为负数")
		return
	}

	if dao.IsIsbnExist(db, isbn) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍已存在")
		return
	}

	book := po.NewBookReference(name, isbn, uint(number), uint(rentNumber))
	dao.CreateBook(db, book)

	response.Response(ctx, http.StatusOK, 200, gin.H{}, "书本录入成功")
}

/*
param:

	isbn
*/
func DeleteBookController(ctx *gin.Context) {
	isbn := ctx.PostForm("isbn")
	err := dao.DeletetBook(dao.GetDb(), isbn)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "不存在的书籍信息")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "删除成功")
}

/*
*
param:

	isbn
	number of addition
*/
func AddBookController(ctx *gin.Context) {
	isbn := ctx.PostForm("isbn")
	addition, err := strconv.Atoi(ctx.PostForm("number"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "isbn必须为数字")
		return
	}
	for {
		book, err := dao.GetBookByIsbn(dao.GetDb(), isbn)
		if err != nil {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
		book.Number += uint(addition)
		uerr := dao.UpdateBook(dao.GetDb(), book)
		if uerr == nil {
			break
		}
		if uerr.Error() == "dml:null value" {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "添加成功")
}

/*
*
param:
isbn
reduceBook
*/
func ReduceBookController(ctx *gin.Context) {
	isbn := ctx.PostForm("isbn")
	reduce, err := strconv.Atoi(ctx.PostForm("number"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "isbn必须为数字")
		return
	}
	for {
		book, err := dao.GetBookByIsbn(dao.GetDb(), isbn)
		if err != nil {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
		if reduce > int(book.Number) {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍存量不能为负")
			return
		}
		book.Number -= uint(reduce)
		uerr := dao.UpdateBook(dao.GetDb(), book)
		if uerr == nil {
			break
		}
		if uerr.Error() == "dml:null value" {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "移除成功")
}

/*
*
param:
isbn
rentNumber
*/
func RentBookController(ctx *gin.Context) {
	isbn := ctx.PostForm("isbn")
	addition, err := strconv.Atoi(ctx.PostForm("number"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "isbn必须为数字")
		return
	}
	for {
		book, err := dao.GetBookByIsbn(dao.GetDb(), isbn)
		if err != nil {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
		book.RentNumber += uint(addition)
		if book.RentNumber > book.Number {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "借出书籍不能大于馆内图书数")
			return
		}
		uerr := dao.UpdateBook(dao.GetDb(), book)
		if uerr == nil {
			break
		}
		if uerr.Error() == "dml:null value" {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "租出成功")
}

/*
*
isbn
returnNumber
*/
func ReturnBookController(ctx *gin.Context) {
	isbn := ctx.PostForm("isbn")
	reduce, err := strconv.Atoi(ctx.PostForm("number"))
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "isbn必须为数字")
		return
	}
	for {
		book, err := dao.GetBookByIsbn(dao.GetDb(), isbn)
		if err != nil {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
		if reduce > int(book.Number) {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "借出书籍不能为负")
			return
		}
		book.RentNumber -= uint(reduce)
		uerr := dao.UpdateBook(dao.GetDb(), book)
		if uerr == nil {
			break
		}
		if uerr.Error() == "dml:null value" {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "书籍不存在")
			return
		}
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{}, "移除成功")
}

func AllBookController(ctx *gin.Context) {
	books := dao.GetAllBook(dao.GetDb())
	if books == nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "查询书籍时发生意外错误")
		return
	}
	bookDtos := dto.NewBookDtosByArray(books)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": bookDtos,
		"count":   len(bookDtos),
	}, "书籍查询成功")
}

func SelectBookController(ctx *gin.Context) {
	isbn := ctx.Param("isbn")
	fmt.Println("isbn:", isbn)
	books := dao.SeleteBook(dao.GetDb(), isbn)
	if books == nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "查询书籍时发生意外错误")
		return
	}
	bookDtos := dto.NewBookDtosByArray(books)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": bookDtos,
		"count":   len(bookDtos),
	}, "书籍查询成功")
}
