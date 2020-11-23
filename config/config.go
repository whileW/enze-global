package config

var conf Config

func init()  {
	//初始化settins
	conf.Setting = Settings{}
}

func InitConfg() *Config {
	//环境变量加载配置
	initEnv(&conf)
	//命令行加载配置
	//initCommand(&conf)
	//配置文件
	initFile(&conf)
	//配置中心
	//设置默认值
	conf.SysSetting.SetDefault()
	return &conf
}
