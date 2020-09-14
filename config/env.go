package config

import "os"

func initEnv(config *Config) {
	config.SysSetting.Env = os.Getenv("ENV")
	config.SysSetting.ConfFilePath = os.Getenv("CFNAME")
	config.SysSetting.HttpAddr = os.Getenv("HTTPADDR")
	config.SysSetting.RpcAddr = os.Getenv("RPCADDR")
	config.SysSetting.Host = os.Getenv("HOST")
}