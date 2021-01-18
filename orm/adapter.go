package orm

import (
	"gorm.io/gorm"
)

type orm_inter interface {
	Open(username,password,path,db_name,config string) gorm.Dialector
	GetConn(username,password,path,db_name,config string) string
}
var OrmInters = map[string]orm_inter{}

func init()  {
	OrmInters["mysql"] = &MysqlAdapter{}
	OrmInters["mssql"] = &MssqlAdapter{}
	OrmInters["sqlserver"] = &MssqlAdapter{}
}
