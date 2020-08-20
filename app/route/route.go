package route

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"github.com/unrolled/secure"
	"net/http"
	"os"
	"strconv"
	"sync"
	"taskdash/app/controller"
	"taskdash/config"
	"taskdash/util"
	"time"
)

func StartHttpServer() error {
	//启动http服务
	addr := config.GetListenAddr()
	fmt.Println("http服务正在启动，监听端口:", util.GetLocalIp()+addr, ",PID:", strconv.Itoa(os.Getpid()))
	r := gin.New()
	SetupHttp(r)
	if config.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	var err error
	go func() {
		if config.GetEnableHttps() {
			r.Use(TlsHandler())
			err = r.RunTLS(addr, "config/cert.pem", "config/key.pem")
		}else {
			server := &http.Server{
				Addr:         addr,
				WriteTimeout: 20 * time.Second,
				Handler:      r,
			}
			err = gracehttp.Serve(server)
		}
		if err != nil {
			return
		}
	}()
	return err
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

func StartNsqServer() error {
	addr := config.GetNsqAddr()
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	controller.ServiceHeartBeat = make(map[string]int64)
	var err error
	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight=9

		err = ServiceRegister(addr, config)
		if err != nil {
			return
		}
		err = TaskRegister(addr, config)
		if err != nil {
			return
		}
		err = UpdateStatus(addr, config)
		if err != nil {
			return
		}
		err = HeartBeat(addr, config)
		if err != nil {
			return
		}
		select{}
	}()
	waiter.Wait()
	return err
}