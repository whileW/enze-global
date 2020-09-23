package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code 	ResCode         `json:"code"`
	Data 	interface{} 	`json:"data"`
	Msg  	interface{}    `json:"msg"`
}

type BasePageModel struct {
	Data		interface{}			`json:"data"`
	Count		int					`json:"count"`
}

type ResCode int
const (
	SUCCESS 		ResCode			= 	0
	//服务器内部逻辑异常
	ERROR   		ResCode			= 	500
	//请求参数检查错误
	ParamterError	ResCode			=	501
	NoFind			ResCode			=	404
	//登陆异常
	LoginFailure	ResCode			=	401
	//鉴权失败
	Unauthorized	ResCode			=	402
)

func Result(c *gin.Context,ResCode ResCode,ResMsg interface{},ResData interface{})  {
	c.JSON(http.StatusOK,Response{
		Code:ResCode,
		Msg:ResMsg,
		Data:ResData,
	})
	c.Abort()
}
func NoFindResult(c *gin.Context)  {
	c.Status(http.StatusNotFound)
	c.Abort()
}
func Ok(c *gin.Context)  {
	Result(c,SUCCESS,"操作成功",nil)
}
func OkWithMessage(c *gin.Context,ResMsg interface{}) {
	Result(c,SUCCESS, ResMsg, nil)
}
func OkWithData(c *gin.Context,ResData interface{}) {
	Result(c,SUCCESS, "操作成功", ResData)
}
func OkDetailed(c *gin.Context,ResMsg interface{},ResData interface{}) {
	Result(c,SUCCESS,ResMsg, ResData)
}

func Fail(c *gin.Context) {
	Result(c,ERROR,"操作失败",nil)
}

func FailWithMessage(c *gin.Context,ResMsg interface{}) {
	Result(c,ERROR, ResMsg,nil)
}
func FailWithData(c *gin.Context,ResData interface{}) {
	Result(c,ERROR, "操作失败",ResData)
}
func FailWithDetailed(c *gin.Context,ResCode ResCode, ResMsg interface{}, ResData interface{}) {
	Result(c,ResCode, ResMsg, ResData)
}
