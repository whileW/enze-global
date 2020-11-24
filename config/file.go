package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
		config.AnalysisSetting(v.AllSettings())
	})
	config.AnalysisSetting(v.AllSettings())
	config.vp = v
}