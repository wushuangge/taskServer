package main

import "C"
import (
	"fmt"
	"log"
	"task/app/route"
	"task/app/store/cache"
	"task/app/store/mongodb"
	"task/app/store/redis"
	"task/config"
	"task/util"
)

func main() {
	//初始化配置文件
	config.InitEnvConf()
	//初始化日志
	util.InitLogger()
	//初始化缓存
	cache.InitMemoryCache()
	//初始化redis
	err := redis.InitRedis()
	if err != nil {
		fmt.Println(err)
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	mongodb.InitMongoDB()
	//开启httpserver
	route.StartHttpServer()
	//开启nsqserver
	route.StartNsqServer()

	select {}
}
