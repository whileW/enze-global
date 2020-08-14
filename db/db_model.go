package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type MySql struct {
	mysqls 			map[string]*gorm.DB		//mysql 链接
}

func NewMySql() *MySql {
	return &MySql{
		mysqls: map[string]*gorm.DB{},
	}
}

func (m *MySql)Get(name string) *gorm.DB {
	if v,ok := m.mysqls[name];ok {
		return v
	}else {
		fmt.Println("没有该MySql实例")
		os.Exit(0)
		return nil
	}
}