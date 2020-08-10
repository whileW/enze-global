package config

import "os"

func initEnv(config *Config) {
	config.SysSetting.Env = os.Getenv("ENV")
	config.SysSetting.ConfFileName = os.Getenv("CFNAME")
}