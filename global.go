package global

import (
	"github.com/whileW/enze-global/config"
	"github.com/whileW/enze-global/log"
	"github.com/whileW/enze-global/mysql"
)

var (
	//配置
	GVA_CONFIG	*config.Config
	//日志
	GVA_LOG		*log.Log
	//Mysql
	GVA_MYSQL 	*mysql.MySql
)

func init() {
	GVA_CONFIG = config.InitConfg()
	GVA_LOG = log.InitLog()
	GVA_MYSQL = mysql.InitMySql(GVA_CONFIG)
}