package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Code int
	Data map[string]interface{}
	Msg  string
}

type ResponseJsons struct {
	Code int
	Data ResponseArrayJson
	Msg  string
}

type ResponseArrayJson struct {
	Payload []map[string]interface{}
	Count   int
}

// union json form
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
