package config

import (
	"github.com/spf13/viper"
)

var conf Config

func init()  {
	//初始化settins
	conf.Setting = settings{}
}

//总配置
type Config struct {
	SysSetting			sysSetting
	Setting				settings
	vp     				*viper.Viper
}

//系统设置
type sysSetting struct {
	//环境 - 默认debug
	//debug - 测试
	//release - 正式
	Env 			string

	//配置文件地址
	ConfFileName 	string
}

//其他设置
type settings map[string]setting
type setting struct {
	Key 		string
	Value 		interface{}
}

func (s *settings)GetInt(key string) int {
	if v, ok := (*s)[key]; ok {
		return v.Value.(int)
	} else {
		panic("key not find")
	}
}
func (s *settings)GetString(key string) string {
	if v, ok := (*s)[key]; ok {
		return v.Value.(string)
	} else {
		panic("key not find")
	}
}
func (s *settings)GetChild(key string) *settings {
	if v, ok := (*s)[key]; ok {
		return v.Value.(*settings)
	} else {
		panic("key not find")
	}
}