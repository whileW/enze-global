package initialize

import (
	"fmt"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/config"
)

func Db()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("初始化db失败：%s\n", r)
			return
		}
	}()
	conf := global.GVA_CONFIG
	//dbs := global.GVA_DB

	db_s,ch := conf.Setting.GetChildd_c("db")
	init_db(db_s)
	go func() {
		for {
			select {
			case <-ch:
				init_db(conf.Setting.GetChildd("db"))
				global.GVA_LOG.Info("监听到数据库配置修改,重新初始化数据库")
			}
		}
	}()
	global.GVA_DB.IsInit = true
}

func init_db(db_s *config.Settings)  {
	for k,_ := range *db_s {
		switch k {
		case "mysql":
			v := db_s.GetChildd("mysql")
			InitMySql(v)
			break
		case "mssql":
			v := db_s.GetChildd("mssql")
			InitMsSql(v)
			break
		}
	}
}