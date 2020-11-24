package initialize

import (
	"github.com/etcd-io/etcd/clientv3"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/etcd"
	"github.com/whileW/enze-global/utils"
	"strings"
	"time"
)

func Etcd()  {
	conf := global.GVA_CONFIG

	v,ch := conf.Setting.GetStringd_c("etcd","")

	init_etcd(v)

	go func() {
		for {
			select {
			case <-ch:
				global.GVA_LOG.Info("监听到etcd配置修改,重新初始化etcd")
				init_etcd(conf.Setting.GetStringd("etcd",""))
			}
		}
	}()
}

func init_etcd(endpoints string)  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoints,";"),
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	global.GVA_ETCD = &etcd.Etcd{Cli:cli,IsInit:true}
}

func RegisterByEtcdRPC(app_name string)  {
	conf := global.GVA_CONFIG
	name := "/rpc/"+conf.SysSetting.Env+"/"+app_name+"/"+utils.RandomString(10)
	register_by_etcd(name,conf.SysSetting.Host+":"+conf.SysSetting.RpcAddr)
}
func RegisterByEtcdHTTP(app_name string)  {
	conf := global.GVA_CONFIG
	name := "/http/"+conf.SysSetting.Env+"/"+app_name+"/"+utils.RandomString(10)
	register_by_etcd(name,conf.SysSetting.Host+":"+conf.SysSetting.HttpAddr)
}

func register_by_etcd(name,host string)  {
	end_ch := make(chan int)
	global.GVA_ETCD.PutLease(name,host,2,end_ch)
	go func(chan<- int) {
		select {
		case <-end_ch:
			panic("etcd lease end")
		}
	}(end_ch)
}