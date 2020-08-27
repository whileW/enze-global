package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type DB struct {
	IsInit			bool					//是否初始化
	dbs 			map[string]*gorm.DB		//db 链接
}

func NewDB() *DB {
	return &DB{
		dbs: map[string]*gorm.DB{},
	}
}

func (m *DB)Set(name string,db *gorm.DB)  {
	m.dbs[name] = db
}

func (m *DB)Get(name string) *gorm.DB {
	if v,ok := m.dbs[name];ok {
		return v
	}else {
		fmt.Println("没有该DB实例")
		os.Exit(0)
		return nil
	}
}