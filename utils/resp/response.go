package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code 	ResCode         `json:"code"`
	Data 	interface{} 	`json:"data"`
	Msg  	string      	`json:"msg"`
}

type ResCode int
const (
	ERROR   		ResCode			= 	500
	SUCCESS 		ResCode			= 	0
	NoFind			ResCode			=	404
	LoginFailure	ResCode			=	401
	Unauthorized	ResCode			=	402
)

func Result(c *gin.Context,ResCode ResCode,ResMsg string,ResData interface{})  {
	c.JSON(http.StatusOK,Response{
		Code:ResCode,
		Msg:ResMsg,
		Data:ResData,
	})
}
func Ok(c *gin.Context)  {
	Result(c,SUCCESS,"操作成功",nil)
}
func OkWithMessage(c *gin.Context,ResMsg string) {
	Result(c,SUCCESS, ResMsg, nil)
}
func OkWithData(c *gin.Context,ResData interface{}) {
	Result(c,SUCCESS, "操作成功", ResData)
}
func OkDetailed(c *gin.Context,ResMsg string,ResData interface{}) {
	Result(c,SUCCESS,ResMsg, ResData)
}

func Fail(c *gin.Context) {
	Result(c,ERROR,"操作失败",nil)
}

func FailWithMessage(c *gin.Context,ResMsg string) {
	Result(c,ERROR, ResMsg,nil)
}
func FailWithData(c *gin.Context,ResData interface{}) {
	Result(c,ERROR, "操作失败",ResData)
}
func FailWithDetailed(c *gin.Context,ResCode ResCode, ResMsg string, ResData interface{}) {
	Result(c,ResCode, ResMsg, ResData)
}
