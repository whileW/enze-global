package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/config"
	"os"
)

//func MySql()  {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Printf("初始化mysql失败：%s\n", r)
//			return
//		}
//	}()
//	conf := global.GVA_CONFIG
//	dbs := global.GVA_DB
//
//	mysql_s := conf.Setting.GetChildd("mysql")
//	if mysql_s == nil {
//		return
//	}
//	for k,_ := range *mysql_s {
//		s := mysql_s.GetChild(k)
//		username,password,path := s.GetStringd("username","root"),s.GetString("password"),s.GetString("path")
//		db_name,config := s.GetString("db_name"),s.GetStringd("config","charset=utf8&parseTime=True&loc=Local")
//		max_idle_conns,max_open_conns := s.GetIntd("max-idle-conns",10),s.GetIntd("max-open-conns",10)
//		log_mode := s.GetBoold("log_mode",true)
//		if db := init_db(username,password,path,db_name,config,max_idle_conns,max_open_conns,log_mode);db != nil{
//			dbs.Set(k,db)
//		}
//	}
//}

func InitMySql(mysql_s *config.Settings) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("初始化mysql失败：%s\n", r)
			return
		}
	}()
	if mysql_s == nil {
		return
	}
	dbs := global.GVA_DB

	for k,_ := range *mysql_s {
		s := mysql_s.GetChild(k)
		username,password,path := s.GetStringd("username","root"),s.GetString("password"),s.GetStringd("path","127.0.0.1:3306")
		db_name,config := s.GetString("db_name"),s.GetStringd("config","charset=utf8&parseTime=True&loc=Local")
		max_idle_conns,max_open_conns := s.GetIntd("max-idle-conns",10),s.GetIntd("max-open-conns",10)
		log_mode := s.GetBoold("log_mode",true)
		if db := init_mysql(username,password,path,db_name,config,max_idle_conns,max_open_conns,log_mode);db != nil{
			dbs.Set(k,db)
		}
	}
}

func init_mysql(username,password,path,db_name,config string,max_idle_conns,max_open_conns int,log_mode bool) *gorm.DB {
	if db, err := gorm.Open("mysql", username+":"+password+"@("+path+")/"+db_name+"?"+config); err != nil {
		fmt.Println("MySQL启动异常:"+err.Error())
		os.Exit(0)
		return nil
	} else {
		db.DB().SetMaxIdleConns(max_idle_conns)
		db.DB().SetMaxOpenConns(max_open_conns)
		db.LogMode(log_mode)
		return db
	}
}