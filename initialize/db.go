package initialize

import (
	"fmt"
	"github.com/whileW/enze-global"
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

	db_s := conf.Setting.GetChildd("db")
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