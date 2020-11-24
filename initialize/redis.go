package initialize

import (
	"github.com/go-redis/redis/v8"
	"github.com/whileW/enze-global"
)

func Redis()  {
	conf := global.GVA_CONFIG

	v,ch := conf.Setting.GetStringd_c("redis","")
	init_redis(v)

	go func() {
		for  {
			select {
			case <- ch:
				global.GVA_LOG.Info("监听到redis配置修改,重新初始化redis")
				init_redis(conf.Setting.GetStringd("redis",""))
			}
		}
	}()
}

func init_redis(v string)  {
	pool := redis.NewClient(&redis.Options{
		Addr:     	v,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	global.GVA_REDIS = pool
}