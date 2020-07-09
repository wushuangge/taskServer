package main

import "C"
import (
	"fmt"
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
	err := store.InitRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	//开启httpserver
	route.StartHttpServer()
	//开启nsqserver
	route.StartNsqServer()

	select {}
}
