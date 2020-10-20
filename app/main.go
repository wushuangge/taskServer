package main

import (
	log "github.com/sirupsen/logrus"
	"taskdash/app/controller"
	"taskdash/app/route"
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
	//初始化定时器
	controller.InitTicker()
	//初始化mongodb
	err = mongodb.InitMongoDB()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	//开启httpserver
	err = route.StartHttpServer()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	route.LoadGroupsInfo()
	//开启nsqserver
	err = route.StartNsqServer()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	select {}
}
