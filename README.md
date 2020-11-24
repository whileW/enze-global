# 项目文档

## 1. 基本介绍
全局配置
### 1.1功能
- 自动读取环境变量
- 自动读取配置文件
- 监听配置文件变化
- 根据配置文件初始化数据库
- 根据配置文件初始化etcd
- 根据配置文件初始化redis
- 日志
- 工具类

## 2. 使用说明
```
go get github.com/whileW/enze-global
```

```
初始化数据库
模仿config.yaml.temp中得数据库配置
在程序入口处调用：initialize.Db()
```

```
配置读取:参考config/*
```

## ETCD
RegisterByEtcdHTTP|RegisterByEtcdRPC
默认注册失效时间为1s