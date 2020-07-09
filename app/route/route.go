package route

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"task/config"
	"task/util"
	"time"
)

func StartHttpServer()  {
	//启动http服务
	log.Println("http服务正在启动，监听端口:", util.GetLocalIp()+":8080", ",PID:", strconv.Itoa(os.Getpid()))
	r := gin.New()
	SetupHttp(r)
	if config.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	go func() {
		var err error
		if runtime.GOOS == "windows" {
			err = http.ListenAndServe(":8080", r)
		} else {
			server := &http.Server{
				Addr:         ":8080",
				WriteTimeout: 20 * time.Second,
				Handler:      r,
			}
			err = gracehttp.Serve(server)
		}

		if err != nil {
			log.Fatal("服务器启动失败:", err.Error())
		}
	}()
}

func StartNsqServer(){
	url := config.GetNsqAddr()
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight=9

		TaskService(url, config)
		HeartBeat(url, config)

		select{}
	}()

	waiter.Wait()
}