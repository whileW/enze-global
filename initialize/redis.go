package initialize

import (
	"github.com/gomodule/redigo/redis"
	"github.com/whileW/enze-global"
)

func Redis()  {
	conf := global.GVA_CONFIG
	pool := redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Setting.GetString("redis"))
		},
	}
	global.GVA_REDIS = &pool
}