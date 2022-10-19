package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wh0rigin/GraduationProject/response"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func ControlController(ctx *gin.Context) {
	apiTag := ctx.PostForm("name")
	//string "0" or "1"
	status := ctx.PostForm("status")

	username := viper.GetString("nlecloud.username")
	password := viper.GetString("nlecloud.password")
	deviceId := viper.GetString("nlecloud.deviceId")
	// Create a Resty Client
	client := resty.New()

	body := fmt.Sprintf(`{"Account":"%s","Password":"%s"}`, username, password)

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://api.nlecloud.com/Users/Login")
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "获取失败")
		return
	}

	var loginResp response.LoginResp
	if err := json.Unmarshal(resp.Body(), &loginResp); err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "Json对象转换失败:"+err.Error())
		return
	}

	token := loginResp.ResultObj.AccessToken
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"deviceId":    deviceId,
			"apiTag":      apiTag,
			"AccessToken": token,
		}).
		SetBody(status).
		SetAuthToken(token).
		Post("http://api.nlecloud.com/Cmds")
	if err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "获取失败")
		return
	}

	var cmdResp response.CmdResp
	if err := json.Unmarshal(resp.Body(), &cmdResp); err != nil {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "Json对象转换失败:"+err.Error())
		return
	}

	if cmdResp.Status != response.Success {
		response.Response(ctx, http.StatusOK, 422, gin.H{}, "控制失败:"+cmdResp.Msg)
		return
	}

	response.Response(ctx, http.StatusOK, 200, gin.H{}, "控制成功")

}
