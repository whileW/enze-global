package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global/utils/resp"
	"testing"
)

type test_req_struct struct {
	Field		string			`json:"field"`
} 

func TestLogGin(t *testing.T)  {
	gin.DefaultWriter = &DisableGinDefaultLog{}
	r := gin.Default()
	r.Use(EnableGinLog())
	r.GET("/", func(c *gin.Context) {
		resp.OkWithMessage(c,"请求成功")
		return
	})
	r.POST("/", func(c *gin.Context) {
		reqBody := &test_req_struct{}
		c.ShouldBindJSON(reqBody)
		fmt.Println(reqBody)
		resp.OkWithMessage(c,"请求成功")
		return
	})
	r.POST("/disable_req_log",DisableGinReqBodyLog(), func(c *gin.Context) {
		reqBody := &test_req_struct{}
		c.ShouldBindJSON(reqBody)
		fmt.Println(reqBody)
		resp.OkWithMessage(c,"请求成功")
		return
	})
	r.Use(DisableGinRespBodyLog())
	r.GET("/disable_log", func(c *gin.Context) {
		resp.OkWithMessage(c,"禁用日志")
		return
	})
	r.Run(":8080")
}