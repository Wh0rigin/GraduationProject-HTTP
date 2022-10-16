package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Code int
	Data map[string]interface{}
	Msg  string
}

// union json form
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
