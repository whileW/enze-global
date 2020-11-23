package main

import (
	"fmt"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
)

func main()  {
	initialize.Etcd()
	//fmt.Println(global.GVA_ETCD.Put("/enze/test1","127.0.0.1"))
	//fmt.Println(global.GVA_ETCD.Put("/enze/test2","127.0.0.2"))
	//global.GVA_ETCD.Delete("/enze/test2")
	initialize.RegisterByEtcdHTTP("test")
	fmt.Println(global.GVA_ETCD.GetWithPrefix(""))
}