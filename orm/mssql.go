package orm

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MssqlAdapter struct{}

func (m *MssqlAdapter)Open(username,password,path,db_name,config string) gorm.Dialector {
	dsn := m.GetConn(username,password,path,db_name,config)
	return sqlserver.Open(dsn)
}
func (m *MssqlAdapter)GetConn(username,password,path,db_name,config string) string {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",username,password,path,db_name)
	return dsn
}