package main

import (
	"encoding/json"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
)

func main()  {
	initialize.MySql()

	file := &File{}
	db := global.GVA_DB.Get("invoices")
	db = db.Table("file").First(file,"id = '0037df68-4875-437f-a37a-f5f1ff766ad3'")
	file_str,_ := json.Marshal(file)
	global.GVA_LOG.Infow("db","result",string(file_str))
}

type File struct {
	Id 					string			`json:"id" gorm:"primary_key;type:varchar(128)"`
	Name				string			`json:"name" gorm:"type:varchar(32)"`
	NewName				string			`json:"new_name" gorm:"type:varchar(128)"`	//新文件名称
	Path 				string			`json:"path" gorm:"type:varchar(128)"`
	FileSize 			int64			`json:"file_size" gorm:"type:int"`			//文件大小
	SaveType 			int				`json:"save_type" gorm:"type:int"`			//保存方式 0本地 1七牛 2支付宝
	State 				int				`json:"state" gorm:"type:int"`			//状态 0临时文件 1正常
}