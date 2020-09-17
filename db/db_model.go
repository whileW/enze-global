package db

import (
	"github.com/jinzhu/gorm"
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
		panic("没有该DB实例")
		return nil
	}
}