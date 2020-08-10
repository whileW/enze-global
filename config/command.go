package config

import (
	"flag"
)

var (
	_env 			string
	_conf_file_path string
)

func initCommand(config *Config) {
	flag.StringVar(&_env, "e", "", "运行环境：debug-测试、release-正式")
	flag.StringVar(&_conf_file_path, "f", "", "配置文件地址")
	flag.Parse()
	if config.SysSetting.Env == "" {
		config.SysSetting.Env = _env
	}
	if config.SysSetting.ConfFileName == "" {
		config.SysSetting.ConfFileName = _conf_file_path
	}
}