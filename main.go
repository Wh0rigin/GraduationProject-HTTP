package main

import (
	dao "github.com/Wh0rigin/GraduationProject/dao"
	gin "github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	defer dao.CloseDb()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
