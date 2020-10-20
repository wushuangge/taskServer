package route

import (
	"encoding/gob"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"github.com/unrolled/secure"
	"net/http"
	"os"
	"strconv"
	_struct "taskdash/app/struct"
	"taskdash/config"
	"taskdash/util"
	"time"
)

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

func StartHttpServer() error {
	//启动http服务
	addr := config.GetListenAddr()
	fmt.Println("http服务正在启动，监听端口:", util.GetLocalIp()+addr, ",PID:", strconv.Itoa(os.Getpid()))
	gob.Register(_struct.User{})
	r := gin.Default()
	store := cookie.NewStore([]byte("mxsecret"))
	r.Use(sessions.Sessions("session", store))
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
			err = r.RunTLS(addr, config.GetCertPem(), config.GetKeyPem())
		} else {
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

func StartNsqServer() error {
	addr := config.GetNsqAddr()
	fmt.Println("nsq服务正在启动，监听端口:", util.GetLocalIp()+addr, ",PID:", strconv.Itoa(os.Getpid()))
	var err error
	go func() {
		config := nsq.NewConfig()
		config.MaxInFlight = 9

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
	}()
	return err
}
