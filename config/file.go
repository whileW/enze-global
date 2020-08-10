package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initFile(config *Config) {
	if config.SysSetting.ConfFileName == "" {
		config.SysSetting.ConfFileName = "config.yaml"
	}
	v := viper.New()
	v.SetConfigFile(config.SysSetting.ConfFileName)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: %s \n", err)
		return
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		analysis_setting(&config.Setting,v.AllSettings(),0)
	})
	analysis_setting(&config.Setting,v.AllSettings(),0)
	config.vp = v
}
//解析配置
func analysis_setting(t *settings,s map[string]interface{},h int)  {
	for k,v := range s {
		ss := setting{
			Key:k,
		}
		if v != nil {
			if d,ok:=v.(map[string]interface{});ok {
				ts := &settings{}
				ss.Value = ts
				analysis_setting(ts,d,h+1)
			}else {
				ss.Value = v
			}
		}
		(*t)[k] = ss
	}
}