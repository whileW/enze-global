package global

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global/log"
	"github.com/whileW/enze-global/utils/resp"
)

func InitGin() *gin.Engine {
	//禁用默认得日志
	gin.DefaultWriter = &log.DisableGinDefaultLog{}
	//修改默认得错误日志
	gin.DefaultErrorWriter = &log.GinErrLog{}
	if GVA_CONFIG.SysSetting.Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	//开启gin
	r := gin.Default()
	// 跨域
	r.Use(resp.Cors())
	//捕获异常
	r.Use(gin.Recovery())
	//开启日志
	r.Use(log.EnableGinLog())
	return r
}