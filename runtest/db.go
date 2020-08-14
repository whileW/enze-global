package main

import (
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
)

func main()  {
	initialize.MySql()
	global.GVA_LOG.Info(global.GVA_DB)
}