package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/casbin/casbin/v2/model"
	"gorm.io/gorm"
)

func InitCasbin(text string,db *gorm.DB) {
	var (
		err error
	)
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(fmt.Sprintf("error: model: %s", err))
	}
	a,err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("error: orm: %s", err))
	}
	GVA_CASBIN,err = casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	GVA_CASBIN.LoadPolicy()
}

