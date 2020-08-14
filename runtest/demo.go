package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/utils/resp"
)

func main()  {
	router := gin.Default()
	router.GET("helloword", func(c *gin.Context) {
		name := c.Query("name")
		global.GVA_LOG.Infow("gin","path","helloword","name",name)
		resp.OkWithMessage(c,"hello"+name)
		return
	})
	router.Run(":"+global.GVA_CONFIG.SysSetting.HttpAddr)
}