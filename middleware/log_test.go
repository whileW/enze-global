package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global/log"
	"testing"
	"time"
)

func TestEnableGinLog(t *testing.T) {
	gin.DefaultWriter = &log.DisableGinDefaultLog{}
	gin.DefaultErrorWriter = &log.GinErrLog{}
	r := gin.Default()
	r.Use(EnableGinLog())

	r.POST("test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"test":"测试",
		})
	})
	r.POST("test_string", func(c *gin.Context) {
		c.String(200,"测试string")
	})
	go r.Run(":8080")
	time.Sleep(60*time.Second)
}