package main

import "C"
import (
	"fmt"
	"log"
	"taskdash/app/route"
	"taskdash/app/store/cache"
	"taskdash/app/store/mongodb"
	"taskdash/config"
	"taskdash/util"
)

func main() {
	var err error
	//初始化配置文件
	config.InitEnvConf()
	//初始化日志
	util.InitLogger()
	//初始化缓存
	cache.InitMemoryCache()
	//初始化mongodb
	err = mongodb.InitMongoDB()
	if err != nil {
		fmt.Println(err)
		log.Fatal("服务器启动失败:", err.Error())
		return
	}

	//开启httpserver
	err = route.StartHttpServer()
	if err != nil {
		fmt.Println(err)
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	//开启nsqserver
	err = route.StartNsqServer()
	if err != nil {
		fmt.Println(err)
		log.Fatal("服务器启动失败:", err.Error())
		return
	}

	select {}
}
