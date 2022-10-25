package router

import (
	"github.com/Wh0rigin/GraduationProject/controller"
	"github.com/Wh0rigin/GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.Engine) {
	//book
	r.POST("/api/book/create", middleware.AuthMiddleware(), controller.CreateBookController)
	r.DELETE("/api/book/delete", middleware.AuthMiddleware(), controller.DeleteBookController)
	r.GET("/api/book/all", middleware.AuthMiddleware(), controller.AllBookController)
	r.GET("/api/book/all/nondto", middleware.MangagerMiddleware(), controller.AllBookNonDto)

	//number
	r.GET("/api/book/number/all", middleware.AuthMiddleware(), controller.GetAllBookNumber)
	r.GET("/api/book/number/rent", middleware.AuthMiddleware(), controller.GetRentBookNumber)
	r.GET("/api/book/number/available", middleware.AuthMiddleware(), controller.GetAvailableBookNumber)

	//use ibsn
	r.POST("/api/book/add", middleware.AuthMiddleware(), controller.AddBookController)
	r.POST("/api/book/reduce", middleware.AuthMiddleware(), controller.ReduceBookController)
	r.POST("/api/book/rent", middleware.AuthMiddleware(), controller.RentBookController)
	r.POST("/api/book/return", middleware.AuthMiddleware(), controller.ReturnBookController)

	//api/book/select/isbn/{isbn}
	r.GET("/api/book/select/isbn/:isbn", middleware.AuthMiddleware(), controller.SelectBookController)
}
