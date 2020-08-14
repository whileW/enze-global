package main

import (
	"encoding/json"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
)

func main()  {
	//初始化mysql
	initialize.Db()
	global.GVA_LOG.Info(global.GVA_DB)

	pi := &process_instance_sql{}
	db := global.GVA_DB.Get("test")
	db = db.Table("process_instance_sql").Find(pi,"id = 'DM_BC-81875fee-01dc-4cda-8b42-c2254e4689da'")
	pi_str,_ := json.Marshal(pi)
	global.GVA_LOG.Infow("db","result",string(pi_str))
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

type process_instance_sql struct {
	Id 					string			`json:"id" gorm:"column:Id"`
	SendUser			string			`json:"send_user" gorm:"column:SendUser"`
	ModuleEId			string			`json:"module_e_id" gorm:"column:ModuleEId"`
	ModuleId			string			`json:"module_id" gorm:"column:ModuleId"`
	ProcessInstanceType	string			`json:"process_instance_type" gorm:"column:ProcessInstanceType"`
}