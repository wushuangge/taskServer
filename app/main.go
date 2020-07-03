package main

import "C"
import (
	"task/app/route"
	"task/app/store"
	"task/config"
	"task/util"
)

func main() {
	//初始化配置文件
	config.InitEnvConf()
	//初始化日志
	util.InitLogger()
	//初始化缓存
	store.InitMemoryCache()
	//初始化redis
	store.InitRedis()
	//开启tcpserver
	route.StartTcpxServer()
	//开启httpserver
	route.StartHttpServer()

	select {}
}
