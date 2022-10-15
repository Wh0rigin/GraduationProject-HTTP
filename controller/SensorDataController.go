package controller

import (
	"net/http"
	"strconv"

	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
)

func AllSensorDataController(ctx *gin.Context) {
	stype := ctx.Param("type")
	datas := dao.GetAllSensorDataByType(dao.GetDb(), stype)
	dataDtos := dto.NewSensorDataDtoByArray(datas)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": dataDtos,
		"count":   len(dataDtos),
	}, "数据查询成功")
}

func RecentSensorDataController(ctx *gin.Context) {
	stype := ctx.Param("type")
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil || limit <= 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "最新数据限制必须为数字且大于0")
		return
	}
	datas := dao.GetRecentSensorDataWithType(dao.GetDb(), stype, limit)
	dataDtos := dto.NewSensorDataDtoByArray(datas)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": dataDtos,
		"count":   len(dataDtos),
	}, "数据查询成功")
}
