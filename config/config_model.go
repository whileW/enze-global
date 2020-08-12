package config

import (
	"github.com/spf13/viper"
)

//总配置
type Config struct {
	SysSetting			sysSetting
	Setting				settings
	vp     				*viper.Viper
}

//系统设置
type sysSetting struct {
	//环境 - 默认debug
	//dev	-	开发
	//debug - 测试
	//release - 正式
	Env 			string
	//配置文件地址 - 默认config.yaml
	ConfFilePath 	string
	//http 监听端口地址 - 默认8080
	HttpAddr		string
	//rpc 监听端口地址 - 默认30010
	RpcAddr 		string
}

//设置默认配置文件地址
func (s *sysSetting)SetDefaultConfFileName() {
	if s.ConfFilePath == "" {
		s.ConfFilePath = "config.yaml"
	}
}
//设置默认环境
func (s *sysSetting)SetDefaultEnv() {
	if s.Env == "" {
		s.Env = "debug"
	}
}
//设置默认http监听地址
func (s *sysSetting)SetDefaultHttpAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = "8080"
	}
}
//设置默认rpc监听地址
func (s *sysSetting)SetDefaultRpcAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = "30010"
	}
}
//设置默认值
func (s *sysSetting)SetDefault()  {
	s.SetDefaultEnv()
	s.SetDefaultHttpAddr()
	s.SetDefaultRpcAddr()
}

//其他设置
type settings map[string]setting
type setting struct {
	Key 		string
	Value 		interface{}
}

func (s *settings)Get(key string) interface{} {
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		panic("key not find")
	}
}
func (s *settings)GetInt(key string) int {
	return s.Get(key).(int)
}
func (s *settings)GetString(key string) string {
	return s.Get(key).(string)
}
func (s *settings)GetBool(key string) bool {
	return s.Get(key).(bool)
}
func (s *settings)GetChild(key string) *settings {
	return s.Get(key).(*settings)
}

func (s *settings)Getd(key string,d interface{})interface{} {
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		return d
	}
}
func (s *settings)GetIntd(key string,d int) int {
	return s.Getd(key,d).(int)
}
func (s *settings)GetStringd(key string,d string) string {
	return s.Getd(key,d).(string)
}
func (s *settings)GetBoold(key string,d bool) bool {
	return s.Getd(key,d).(bool)
}
func (s *settings)GetChildd(key string) *settings {
	v := s.Getd(key,nil)
	if v == nil {
		return nil
	}
	return v.(*settings)
}