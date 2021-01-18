package middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestCors(t *testing.T) {
	r := gin.Default()
	r.Use(Cors())
	r.POST("test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"test":"测试",
		})
	})
	go r.Run(":8080")
	time.Sleep(30*time.Second)
}