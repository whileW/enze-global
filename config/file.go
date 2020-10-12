package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

func initFile(config *Config) {
	if config.SysSetting.ConfFilePath == "" {
		config.SysSetting.SetDefaultConfFileName()
	}
	v := viper.New()
	v.SetConfigFile(config.SysSetting.ConfFilePath)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Sprintf("Fatal error config file: %s \n", err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		analysis_setting(config,&config.Setting,v.AllSettings(),0)
	})
	analysis_setting(config,&config.Setting,v.AllSettings(),0)
	config.vp = v
}
//解析配置
func analysis_setting(config *Config,t *Settings,s map[string]interface{},h int)  {
	for k,v := range s {
		if h == 0 {
			set_sys_setting(config,k,v)
		}
		ss := Setting{
			Key:k,
		}
		if v != nil {
			if d,ok:=v.(map[string]interface{});ok {
				ts := &Settings{}
				ss.Value = ts
				analysis_setting(config,ts,d,h+1)
			}else {
				ss.Value = v
			}
		}
		(*t)[k] = ss
	}
}
func set_sys_setting(config *Config,k string,v interface{})  {
	uk := strings.ToUpper(k)
	//val := v.(string)
	switch uk {
	case "ENV":
		if val,ok := v.(string);ok {
			config.SysSetting.Env = val
		}
		break
	case "HTTPADDR":
		if val,ok := v.(int);ok {
			config.SysSetting.Env = strconv.Itoa(val)
		}
		break
	case "RPCADDR":
		if val,ok := v.(int);ok {
			config.SysSetting.RpcAddr = strconv.Itoa(val)
		}
		break
	case "HOST":
		if val,ok := v.(string);ok {
			config.SysSetting.Host = val
		}
		break
	}
}