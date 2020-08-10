package global

import (
	"github.com/whileW/enze-global/config"
)

var (
	//配置
	GVA_CONFIG	*config.Config
	//Mysql
)

func init()  {
	GVA_CONFIG = config.InitConfg()
}