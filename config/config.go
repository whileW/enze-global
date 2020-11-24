package config

import (
	"fmt"
	"strconv"
	"strings"
)

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

//解析配置
func (c *Config)AnalysisSetting(s map[string]interface{})  {
	temp := &Settings{}
	c.analysisSetting(&c.Setting,temp,s,0)
	c.Setting = *temp
	fmt.Println(c.Setting)
}
func (c *Config)analysisSetting(source *Settings,temp *Settings,s map[string]interface{},h int) bool {
	is_change := false
	for k,v := range s {
		if h == 0 {
			c.SetSysSetting(k,v)
		}

		//变量初始化
		var source_s *Setting
		if t,ok := (*source)[k];ok {
			source_s = &t
		}else {
			source_s = &Setting{
				Key:k,
				Value:&Settings{},
			}
		}
		temp_s := &Setting{
			Key:k,
		}

		//配置文件赋值
		if v != nil {
			if d,ok:=v.(map[string]interface{});ok {
				ts := &Settings{}
				temp_s.Value = ts
				if ic := c.analysisSetting(source_s.Value.(*Settings),ts,d,h+1);ic{
					is_change = ic
					source_s.send_change()
				}
			}else {
				if v != source_s.Value {
					is_change = true
					source_s.send_change()
				}
				temp_s.Value = v
			}
		}

		if source_s != nil {
			temp_s.change_chan = source_s.change_chan
		}
		(*temp)[k] = *temp_s
	}
	return is_change
}
//设置系统配置
func (c *Config)SetSysSetting(k string,v interface{})  {
	uk := strings.ToUpper(k)
	//val := v.(string)
	switch uk {
	case "ENV":
		if val,ok := v.(string);ok {
			c.SysSetting.Env = val
		}
		break
	case "HTTPADDR":
		if val,ok := v.(int);ok {
			c.SysSetting.HttpAddr = strconv.Itoa(val)
		}
		break
	case "RPCADDR":
		if val,ok := v.(int);ok {
			c.SysSetting.RpcAddr = strconv.Itoa(val)
		}
		break
	case "HOST":
		if val,ok := v.(string);ok {
			c.SysSetting.Host = val
		}
		break
	}
}