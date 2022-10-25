package controller

import (
	"net/http"
	"strconv"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
)

// AllSensor godoc
// @Summary      Get all Sensor Data
// @Description  get all sensor data by list
// @Tags         /api/sensor/
// @Accept       json
// @Produce      json
// @Param        type     path    string  true  "temperature"
// @Success      200  {object} response.ResponseJsons
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/sensor/all/{type} [GET]
func AllSensorDataController(ctx *gin.Context) {
	stype := ctx.Param("type")
	datas := dao.GetAllSensorDataByType(common.GetDb(), stype)
	dataDtos := dto.NewSensorDataDtoByArray(datas)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": dataDtos,
		"count":   len(dataDtos),
	}, "数据查询成功")
}

// RecentSensor godoc
// @Summary      Get recent Sensor Data
// @Description  get recent sensor data by list
// @Tags         /api/sensor/
// @Accept       json
// @Produce      json
// @Param        type     path    string  true  "temperature"
// @Param        limit     path    int  true   "Number"  Format(int64)
// @Success      200  {object} response.ResponseJsons
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/sensor/recent/{type}/{limit} [GET]
func RecentSensorDataController(ctx *gin.Context) {
	stype := ctx.Param("type")
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil || limit <= 0 {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "最新数据限制必须为数字且大于0")
		return
	}
	datas := dao.GetRecentSensorDataWithType(common.GetDb(), stype, limit)
	dataDtos := dto.NewSensorDataDtoByArray(datas)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": dataDtos,
		"count":   len(dataDtos),
	}, "数据查询成功")
}
