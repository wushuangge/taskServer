package route

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/fwhezfwhez/tcpx"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"task/config"
	"task/util"
	"time"
)

func StartTcpxServer() {
	srv := tcpx.NewTcpX(nil)
	srv.BeforeExit(func() {
		fmt.Println("server stops")
	})

	srv.OnClose = OnClose
	srv.OnConnect = OnConnect
	//srv.HeartBeatModeDetail(true, 10*time.Second, false, tcpx.DEFAULT_HEARTBEAT_MESSAGEID)
	//心跳
	srv.HeartBeatMode(true, 5*time.Second)
	srv.RewriteHeartBeatHandler(tcpx.DEFAULT_HEARTBEAT_MESSAGEID, OnHeartBeat)
	//用户池
	srv.WithBuiltInPool(true)
	//中间件
	srv.UseGlobal(countRequestTime)
	srv.AddHandler(1, getRequestTime)
	srv.AddHandler(taskHandlerMessageID, taskHandler)

	if config.IsDev() {
		tcpx.SetLogMode(tcpx.DEBUG)
	} else {
		tcpx.SetLogMode(tcpx.RELEASE)
	}

	log.Println("tcp服务正在启动，监听端口:", util.GetLocalIp()+":8090", ",PID:", strconv.Itoa(os.Getpid()))
	// tcp
	go func() {
		log.Println("tcp服务正在启动，监听端口:", util.GetLocalIp()+":8090", ",PID:", strconv.Itoa(os.Getpid()))
		if e := srv.ListenAndServe("tcp", ":8090"); e != nil {
			panic(e)
		}
	}()
}

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
