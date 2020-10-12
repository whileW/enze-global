package config

import (
	"flag"
)

var (
	_env 			string
	_conf_file_path string
	_http_addr 		string
	_rpc_addr		string
	_host			string
)

func initCommand(config *Config) {
	flag.StringVar(&_env, "e", "", "运行环境：debug-测试、release-正式")
	flag.StringVar(&_conf_file_path, "f", "", "配置文件地址")
	flag.StringVar(&_http_addr, "ha", "", "http监听地址")
	flag.StringVar(&_rpc_addr, "ra", "", "rpc监听地址")
	flag.StringVar(&_host,"h","","本机ip地址")
	if !flag.Parsed() {
		flag.Parse()
	}
	if _env != "" {
		config.SysSetting.Env = _env
	}
	if _conf_file_path != "" {
		config.SysSetting.ConfFilePath = _conf_file_path
	}
	if _http_addr != "" {
		config.SysSetting.HttpAddr = _http_addr
	}
	if _rpc_addr != "" {
		config.SysSetting.RpcAddr = _rpc_addr
	}
	if _host != "" {
		config.SysSetting.Host = _host
	}
}