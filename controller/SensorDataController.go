package controller

import (
	"net/http"
	"strconv"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/dao"
	"github.com/Wh0rigin/GraduationProject/dto"
	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/Wh0rigin/GraduationProject/util"
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

// GetAllSensorDataNonDto godoc
// @Summary      Get all Sensor Data nonlimit nonDto
// @Description  get all sensor data nonlimit
// @Tags         /api/sensor/
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJsons
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/sensor/all/non/filtrate [GET]
func GetAllSensorDataNonFiltrate(ctx *gin.Context) {
	datas := dao.GetAllSensorData(common.GetDb())
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"payload": datas,
		"count":   len(datas),
	}, "数据查询成功")
}

// GetAllSensorDataLimit godoc
// @Summary      Get all Sensor Data limit
// @Description  get all sensor data limit
// @Tags         /api/sensor/
// @Accept       json
// @Produce      json
// @Success      200  {object} response.ResponseJsons
// @Failure      442  {object} response.ResponseJson
// @Failure      500  {object} response.ResponseJson
// @Router       /api/sensor/all/page [GET]
func GetAllSensorDataLimit(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "限制数量错误")
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "分页错误")
	}
	sort := ctx.Query("sort")
	stype := ctx.Query("type")
	if stype == "" {
		datas := dao.GetAllSensorData(common.GetDb())
		if sort == "-id" {
			util.ReverseSensorData(datas)
		}
		if limit*page < len(datas) {
			response.Response(ctx, http.StatusOK, 200, gin.H{
				"payload": datas[limit*(page-1) : limit*page],
				"count":   len(datas),
			}, "数据查询成功")
		} else {
			response.Response(ctx, http.StatusOK, 200, gin.H{
				"payload": datas[limit*(page-1):],
				"count":   len(datas),
			}, "数据查询成功")
		}
		return
	}
	datas := dao.GetAllSensorDataByType(common.GetDb(), stype)
	if sort == "-id" {
		util.ReverseSensorData(datas)
	}
	if limit*page < len(datas) {
		response.Response(ctx, http.StatusOK, 200, gin.H{
			"payload": datas[limit*(page-1) : limit*page],
			"count":   len(datas),
		}, "数据查询成功")
	} else {
		response.Response(ctx, http.StatusOK, 200, gin.H{
			"payload": datas[limit*(page-1):],
			"count":   len(datas),
		}, "数据查询成功")
	}
}

func SensorDataNumberByName(ctx *gin.Context) {
	sensorName := ctx.Param("sensorName")
	datas := dao.GetAllSensorDataByType(common.GetDb(), sensorName)
	response.Response(ctx, http.StatusOK, 200, gin.H{
		"total": len(datas),
	}, "数据查询成功")
}
