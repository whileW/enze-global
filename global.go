package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/whileW/enze-global/config"
	"github.com/whileW/enze-global/etcd"
	"github.com/whileW/enze-global/log"
	"github.com/whileW/enze-global/orm"
)

var (
	//配置
	GVA_CONFIG	*config.Config
	//日志
	GVA_LOG		*log.Loger
	//Db
	GVA_DB 		*orm.Orm
	//ETCD
	GVA_ETCD	*etcd.Etcd
	//REDIS
	GVA_REDIS	*redis.Client
	//casbin
	GVA_CASBIN 	*casbin.Enforcer
)

func init() {
	GVA_CONFIG = config.InitConfg()
	GVA_LOG = log.GetLoger()
	GVA_DB = orm.GetOrm()
}

func IsHaveRedis() bool {
	if GVA_CONFIG.Setting.GetStringd("redis","") != ""{
		return true
	}
	return false
}
func IsHaveETCD() bool {
	if GVA_CONFIG.Setting.GetStringd("etcd","") != ""{
		return true
	}
	return false
}