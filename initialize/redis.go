package initialize

import (
	"github.com/go-redis/redis/v8"
	"github.com/whileW/enze-global"
)

func Redis()  {
	conf := global.GVA_CONFIG
	pool := redis.NewClient(&redis.Options{
		Addr:     	conf.Setting.GetString("redis"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	global.GVA_REDIS = pool
}