package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlAdapter struct{}

func (m *MysqlAdapter)Open(username,password,path,db_name,config string) gorm.Dialector {
	dsn := m.GetConn(username,password,path,db_name,config)
	return mysql.Open(dsn)
}
func (m *MysqlAdapter)GetConn(username,password,path,db_name,config string) string {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s",username,password,path,db_name,config)
	return dsn
}