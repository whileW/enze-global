package initialize

import (
	"github.com/etcd-io/etcd/clientv3"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/etcd"
	"strings"
	"time"
)

func Etcd()  {
	conf := global.GVA_CONFIG
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(conf.Setting.GetString("etcd"),";"),
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	global.GVA_ETCD = &etcd.Etcd{Cli:cli,IsInit:true}
}

func RegisterByEtcd(app_name string)  {
	conf := global.GVA_CONFIG
	end_ch := make(chan int)
	global.GVA_ETCD.PutLease(app_name,conf.SysSetting.Host,500,end_ch)
	go func(chan<- int) {
		select {
		case <-end_ch:
			panic("etcd lease end")
		}
	}(end_ch)
}