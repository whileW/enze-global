package orm

import (
	"fmt"
	"github.com/whileW/enze-global/config"
	"github.com/whileW/enze-global/log"
	"gorm.io/gorm"
	"os"
)

//orm 依赖conf和log

func init()  {
	InitOrm(config.GetConf())
}

func InitOrm(conf *config.Config) *Orm {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("初始化db失败：%s\n", r)
	//		return
	//	}
	//}()

	db_s,ch := conf.Setting.GetChildd_c("db")
	orm = NewOrm()
	init_orms(db_s,orm)
	go func() {
		for {
			select {
			case <-ch:
				init_orms(conf.Setting.GetChildd("db"),orm)
				log.GetLoger().Info("监听到数据库配置修改,重新初始化数据库")
			}
		}
	}()
	return orm
}

func init_orms(db_s *config.Settings,orms *Orm)  {
	for k,_ := range *db_s {
		v := db_s.GetChildd(k)
		adapter := OrmInters[k]
		for kk,_ := range *v {
			s := v.GetChild(kk)
			orm_db := init_orm(s,adapter)
			orms.Set(kk,orm_db)
		}
	}
}
func init_orm(s *config.Settings,adapter orm_inter) *gorm.DB {
	username,password,path := s.GetStringd("username","root"),s.GetString("password"),s.GetStringd("path","127.0.0.1:3306")
	db_name,config := s.GetString("db_name"),s.GetStringd("config","charset=utf8&parseTime=True&loc=Local")
	max_idle_conns,max_open_conns := s.GetIntd("max-idle-conns",10),s.GetIntd("max-open-conns",10)
	oc := &gorm.Config{
		Logger:&log.GormLogger{},
	}
	if db, err := gorm.Open(adapter.Open(username,password,path,db_name,config),oc); err != nil {
		fmt.Println("MySQL启动异常:"+err.Error())
		os.Exit(0)
		return nil
	} else {
		sql_db,_ := db.DB()
		sql_db.SetMaxIdleConns(max_idle_conns)
		sql_db.SetMaxOpenConns(max_open_conns)
		return db
	}
}