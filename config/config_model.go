package config

import (
	"fmt"
	"github.com/spf13/viper"
)

//总配置
type Config struct {
	SysSetting			sysSetting
	Setting				Settings
	vp     				*viper.Viper
}

var(
	default_conf_file_path = "config.yaml"
	default_env = "debug"
	default_http_addr = "8080"
	default_rpc_addr = "30010"
	default_host = "127.0.0.1"
)

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
	//本机ip
	Host 			string
}

//设置默认配置文件地址
func SetDefaultConfFileName(conf_path string)  {
	default_conf_file_path = conf_path
}
func (s *sysSetting)SetDefaultConfFileName() {
	if s.ConfFilePath == "" {
		s.ConfFilePath = default_conf_file_path
	}
}
//设置默认环境
func (s *sysSetting)SetDefaultEnv() {
	if s.Env == "" {
		s.Env = default_env
	}
}
//设置默认http监听地址
func (s *sysSetting)SetDefaultHttpAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = default_http_addr
	}
}
//设置默认rpc监听地址
func (s *sysSetting)SetDefaultRpcAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = default_rpc_addr
	}
}
//设置默认host
func (s *sysSetting)SetDefaultHost()  {
	if s.Host == "" {
		s.Host = default_host
	}
}
//设置默认值
func (s *sysSetting)SetDefault()  {
	s.SetDefaultEnv()
	s.SetDefaultHttpAddr()
	s.SetDefaultRpcAddr()
	s.SetDefaultHost()
}

//其他设置
type Settings map[string]Setting
type Setting struct {
	Key 			string
	Value 			interface{}
	change_chan 	[]chan int
}

func (s *Settings)Get(key string) interface{} {
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		panic("key not find")
	}
}
func (s *Settings)GetInt(key string) int {
	return s.Get(key).(int)
}
func (s *Settings)GetString(key string) string {
	return s.Get(key).(string)
}
func (s *Settings)GetBool(key string) bool {
	return s.Get(key).(bool)
}
func (s *Settings)GetChild(key string) *Settings {
	return s.Get(key).(*Settings)
}

func (s *Settings)Getd(key string,d interface{})interface{} {
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		return d
	}
}
func (s *Settings)GetIntd(key string,d int) int {
	return s.Getd(key,d).(int)
}
func (s *Settings)GetStringd(key string,d string) string {
	return s.Getd(key,d).(string)
}
func (s *Settings)GetBoold(key string,d bool) bool {
	return s.Getd(key,d).(bool)
}
func (s *Settings)GetChildd(key string) *Settings {
	v := s.Getd(key,nil)
	if v == nil {
		return nil
	}
	return v.(*Settings)
}

func (s *Settings)Getd_c(key string,d interface{})(interface{},chan int) {
	var value interface{}
	change_chan := make(chan int)
	if v, ok := (*s)[key]; ok {
		value = v.Value
		if change_chan == nil || len(v.change_chan) <= 0 {
			v.change_chan = []chan int{}
		}
		v.change_chan = append(v.change_chan, change_chan)
		(*s)[key] = v
	} else {
		value = d
	}
	return value,change_chan
}
func (s *Settings)GetIntd_c(key string,d int) (int,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(int),ch
}
func (s *Settings)GetStringd_c(key string,d string) (string,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(string),ch
}
func (s *Settings)GetBoold_c(key string,d bool) (bool,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(bool),ch
}
func (s *Settings)GetChildd_c(key string) (*Settings,chan int) {
	v,ch := s.Getd_c(key,nil)
	if v == nil {
		return nil,ch
	}
	return v.(*Settings),ch
}

func (s *Setting)send_change()  {
	if s.change_chan != nil {
		for _,t := range s.change_chan {
			fmt.Println("配置修改:",s.Key)
			t<-1
		}
	}
}