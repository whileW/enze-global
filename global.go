package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/whileW/enze-global/config"
	"github.com/whileW/enze-global/db"
	"github.com/whileW/enze-global/etcd"
	"github.com/whileW/enze-global/log"
)

var (
	//配置
	GVA_CONFIG	*config.Config
	//日志
	GVA_LOG		*log.Log
	//Db
	GVA_DB 		*db.DB
	//ETCD
	GVA_ETCD	*etcd.Etcd
	//REDIS
	GVA_REDIS	*redis.Client
)

func init() {
	GVA_CONFIG = config.InitConfg()
	GVA_LOG = log.InitLog()
	GVA_DB = db.NewDB()
}