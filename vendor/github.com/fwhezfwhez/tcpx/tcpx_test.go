package tcpx

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	// "github.com/xtaci/kcp-go"
	"net"
	"runtime"
	"testing"
	"time"
)

func TestTcpX_ListenAndServe(t *testing.T) {
	var onConnect = func(c *Context) {
		fmt.Println(fmt.Sprintf("connecting from remote host %s network %s", c.Conn.RemoteAddr().String(), c.Conn.RemoteAddr().Network()))

	}
	var onClose = func(c *Context) {
		fmt.Println(fmt.Sprintf("connecting from remote host %s network %s has stoped", c.Conn.RemoteAddr().String(), c.Conn.RemoteAddr().Network()))

	}
	var sayHello_1 = func(c *Context) {
		fmt.Println("hello")

	}
	var sayGoodBye_2 = func(c *Context) {
		fmt.Println("good bye")

	}
	tcpx := NewTcpX(JsonMarshaller{})
	tcpx.OnConnect = onConnect
	tcpx.OnClose = onClose
	tcpx.AddHandler(1, sayHello_1)
	tcpx.AddHandler(2, sayGoodBye_2)

	fmt.Println("开始监听: tcp 7676")
	go func() {
		e := tcpx.ListenAndServe("tcp", ":7676")
		if e != nil {
			fmt.Println(e.Error())
			return
		}
	}()
	time.Sleep(10 * time.Second)
}

// Including:
// Usages of global,anchor,router type middleware, messageID 1 shows middlewareOrder [1,2,3]
func TestTcpX_TCP_Middleware(t *testing.T) {
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 1)
	// middlewareOrder suggest the execute order of three kinds middleware [1,2,3]
	var middlewareOrder = make([]int, 0, 10)
	// client
	go func() {
		<-serverStart

		conn, err := net.Dial("tcp", "localhost:7004")
		if err != nil {
			testResult <- errorx.Wrap(err)
			fmt.Println(errorx.Wrap(err).Error())
			return
		}

		buf, e := PackJSON.Pack(1, "hello, I'm client")

		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		conn.Write(buf)
	}()

	// server
	go func() {
		srv := NewTcpX(JsonMarshaller{})
		srv.OnMessage = nil
		srv.BeforeExit(func() {
			fmt.Println("exit")
		})
		// global middleware
		srv.UseGlobal(func(c *Context) {
			middlewareOrder = append(middlewareOrder, 1)
		})
		// anchor middleware
		srv.Use("anchor1", func(c *Context) {
			middlewareOrder = append(middlewareOrder, 2)
		})
		// router middleware
		srv.AddHandler(1, func(c *Context) {
			middlewareOrder = append(middlewareOrder, 3)
		}, func(c *Context) {
			fmt.Println(middlewareOrder)
			if len(middlewareOrder) != 3 {
				testResult <- errorx.NewFromStringf("middlewareOrder len want 3 but got %d", len(middlewareOrder))
				return
			}
			testResult <- nil
			c.Reply(10086, "hello, I'm server")
		})

		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()
		e := srv.ListenAndServeTCP("tcp", ":7004")
		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(e.Error())
			return
		}
	}()

	e := <-testResult
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
	}
}

// including:
// Usage of UnUse, get middlewareOrder [1, 3], 2 is jumped by UnUse
func TestTcpX_UDP_Middleware_UnUse(t *testing.T) {
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 1)
	// middlewareOrder suggest the execute order of three kinds middleware [1,2,3]
	var middlewareOrder = make([]int, 0, 10)
	// client
	go func() {
		<-serverStart

		conn, err := net.Dial("udp", "localhost:7005")
		if err != nil {
			testResult <- errorx.Wrap(err)
			fmt.Println(errorx.Wrap(err).Error())
			return
		}

		buf, e := PackJSON.Pack(2, "hello, I'm client")

		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		conn.Write(buf)
	}()

	// server
	go func() {
		srv := NewTcpX(JsonMarshaller{})
		srv.OnMessage = nil

		// global middleware
		srv.UseGlobal(func(c *Context) {
			middlewareOrder = append(middlewareOrder, 1)
		})
		// anchor middleware
		srv.Use("anchor1", func(c *Context) {
			middlewareOrder = append(middlewareOrder, 2)
		})
		//// router middleware
		//srv.AddHandler(1, func(c *Context) {
		//	middlewareOrder = append(middlewareOrder, 3)
		//}, func(c *Context) {
		//	c.Reply(10086, "hello, I'm server")
		//})

		srv.UnUse("anchor1")
		srv.AddHandler(2, func(c *Context) {
			middlewareOrder = append(middlewareOrder, 3)
		}, func(c *Context) {
			fmt.Println(middlewareOrder)
			if len(middlewareOrder) != 2 {
				testResult <- errorx.NewFromStringf("middlewareOrder len want 2 but got %d, %v", len(middlewareOrder), middlewareOrder)
				return
			}
			testResult <- nil
			c.Reply(10086, "hello, I'm server")
		})
		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()
		e := srv.ListenAndServeUDP("udp", ":7005")
		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(e.Error())
			return
		}
	}()

	e := <-testResult
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
	}
}

// Usage of Abort and Next, get middlewareOrder [1,2,3] 4 is aborted
//func TestTcpX_KCP_Middleware_Abort_Next(t *testing.T) {
//	var serverStart = make(chan int, 1)
//	var testResult = make(chan error, 1)
//	// middlewareOrder suggest the execute order of three kinds middleware [1,2,3]
//	var middlewareOrder = make([]int, 0, 10)
//	// client
//	go func() {
//		<-serverStart
//
//		conn, err := kcp.DialWithOptions("localhost:7006", nil, 10, 3)
//		if err != nil {
//			testResult <- errorx.Wrap(err)
//			fmt.Println(errorx.Wrap(err).Error())
//			return
//		}
//
//		buf, e := PackJSON.Pack(1, "hello, I'm client")
//
//		if e != nil {
//			testResult <- errorx.Wrap(e)
//			fmt.Println(errorx.Wrap(e).Error())
//			return
//		}
//		conn.Write(buf)
//	}()
//
//	// server
//	go func() {
//		srv := NewTcpX(JsonMarshaller{})
//		srv.OnMessage = nil
//
//		// global middleware
//		srv.UseGlobal(func(c *Context) {
//			middlewareOrder = append(middlewareOrder, 1)
//			fmt.Println("pass global")
//		})
//		// anchor middleware
//		srv.Use("anchor1", func(c *Context) {
//			middlewareOrder = append(middlewareOrder, 2)
//			fmt.Println("pass anchor1")
//			c.Next()
//		}, "anchor2", func(c *Context) {
//			middlewareOrder = append(middlewareOrder, 3)
//			fmt.Println("pass anchor2")
//			c.Abort()
//			time.Sleep(2 * time.Second)
//			fmt.Println(middlewareOrder)
//			if len(middlewareOrder) != 3 {
//				testResult <- errorx.NewFromStringf("middlewareOrder len want 3 but got %d", len(middlewareOrder))
//				return
//			}
//			testResult <- nil
//		}, "anchor3", func(c *Context) {
//			fmt.Println("should not pass anchor 3, but passed")
//			middlewareOrder = append(middlewareOrder, 4)
//		})
//
//		// router middleware
//		// no chance to exec since anchor abort the chain
//		srv.AddHandler(1, func(c *Context) {
//		})
//
//		go func() {
//			time.Sleep(time.Second * 10)
//			serverStart <- 1
//		}()
//		e := srv.ListenAndServe("kcp", ":7006")
//		if e != nil {
//			testResult <- errorx.Wrap(e)
//			fmt.Println(e.Error())
//			return
//		}
//	}()
//
//	e := <-testResult
//	if e != nil {
//		fmt.Println(e.Error())
//		t.Fail()
//	}
//}

func TestTcpX_OnMessage(t *testing.T) {
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 1)
	// middlewareOrder suggest the execute order of three kinds middleware [1,2,3]
	var middlewareOrder = make([]int, 0, 10)
	// client
	go func() {
		<-serverStart

		conn, err := net.Dial("tcp", "localhost:7007")
		if err != nil {
			testResult <- errorx.Wrap(err)
			fmt.Println(errorx.Wrap(err).Error())
			return
		}

		buf, e := PackJSON.Pack(1, "hello, I'm client")

		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		conn.Write(buf)
	}()

	// server
	go func() {
		srv := NewTcpX(JsonMarshaller{})
		srv.OnMessage = func(c *Context) {
			fmt.Println(c.Stream)
			bodyBytes, e := srv.Packx.BodyBytesOf(c.Stream)
			if e != nil {
				fmt.Println(errorx.Wrap(e).Error())
				testResult <- errorx.Wrap(e)
				return
			}
			var receive string
			e = json.Unmarshal(bodyBytes, &receive)
			if e != nil {
				fmt.Println(errorx.Wrap(e).Error())
				testResult <- errorx.Wrap(e)
				return
			}
			if receive != "hello, I'm client" {
				testResult <- errorx.NewFromStringf("received want %s but got %s", "hello, I'm client", receive)
				return
			}
			testResult <- nil
		}
		srv.BeforeExit(func() {
			fmt.Println("exit")
		})
		// global middleware
		srv.UseGlobal(func(c *Context) {
			middlewareOrder = append(middlewareOrder, 1)
		})
		// anchor middleware
		srv.Use("anchor1", func(c *Context) {
			middlewareOrder = append(middlewareOrder, 2)
		})

		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()
		e := srv.ListenAndServeTCP("tcp", ":7007")
		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(e.Error())
			return
		}
	}()

	e := <-testResult
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
	}
}

func TestHeartbeat(t *testing.T) {
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 1)
	go func() {
		time.Sleep(40 * time.Second)
		testResult <- nil
	}()

	// client
	go func() {
		<-serverStart

		conn, e := net.Dial("tcp", "localhost:7008")

		if e != nil {
			testResult <- errorx.Wrap(e)
			panic(e)
		}
		var heartBeat []byte
		heartBeat, e = PackWithMarshaller(Message{
			MessageID: DEFAULT_HEARTBEAT_MESSAGEID,
			Header:    nil,
			Body:      nil,
		}, nil)
		for {
			_, e = conn.Write(heartBeat)
			if e != nil {
				fmt.Println(e.Error())
				testResult <- errorx.Wrap(e)
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// server
	go func() {
		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()

		srv := NewTcpX(nil)

		srv.HeartBeatModeDetail(true, 5*time.Second, false, DEFAULT_HEARTBEAT_MESSAGEID)

		//srv.RewriteHeartBeatHandler(1300, func(c *tcpx.Context) {
		//	fmt.Println("rewrite heartbeat handler")
		//	c.RecvHeartBeat()
		//})

		SetLogMode(DEBUG)

		srv.ListenAndServe("tcp", ":7008")
	}()

	e := <-testResult
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
	}
}

func TestHandleRaw(t *testing.T) {
	var srvStart = make(chan int, 1)

	go func() {
		time.Sleep(4 * time.Second)
		srvStart <- 1
	}()
	go func() {
		srv := NewTcpX(nil)
		srv.UseGlobal(func(c *Context) {
			fmt.Println("before raw message in")
		})
		srv.Use("middle-1", func(c *Context) {
			fmt.Println("use middleware 1")
		})
		srv.HandleRaw = func(c *Context) {
			var buf = make([]byte, 500)
			var n int
			var e error
			for {
				n, e = c.ConnReader.Read(buf)
				if e != nil {
					fmt.Println(e.Error())
					return
				}
				fmt.Println("receive:", buf[:n])
				c.ConnWriter.Write([]byte("hello,I am server."))
			}
		}
		srv.ListenAndServeRaw("tcp", ":6631")
	}()
	<-srvStart
	conn, e := net.Dial("tcp", "localhost:6631")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	conn.Write([]byte("hello,I am client."))

	time.Sleep(10 * time.Second)
}

func TestConCurrentTCP(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 100)

	// 100 client  100 messages
	go func() {
		<-serverStart

		for i := 0; i < 150; i++ {
			conn, err := net.Dial("tcp", "localhost:6632")
			if err != nil {
				testResult <- errorx.Wrap(err)
				fmt.Println(errorx.Wrap(err).Error())
				return
			}

			buf, e := PackJSON.Pack(2, fmt.Sprintf("hello, I'm client %d call 2", i))
			if e != nil {
				testResult <- errorx.Wrap(e)
				fmt.Println(errorx.Wrap(e).Error())
				return
			}
			buf2, e2 := PackJSON.Pack(1, "hello, I'm client call 1")
			if e2 != nil {
				testResult <- errorx.Wrap(e)
				fmt.Println(errorx.Wrap(e).Error())
				return
			}
			for j := 0; j < 200; j++ {
				go func(j int) {
					conn.Write(buf)
					conn.Write(buf2)
				}(j)
			}
		}

	}()

	// server
	go func() {
		srv := NewTcpX(JsonMarshaller{})

		srv.AddHandler(2, func(c *Context) {
			time.Sleep(15 * time.Second)
			c.Reply(10086, fmt.Sprintf("hello, I'm server"))
		})
		srv.AddHandler(1, func(c *Context) {
			time.Sleep(15 * time.Second)
			c.Reply(10086, fmt.Sprintf("hello, I'm server"))
		})
		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()
		e := srv.ListenAndServe("tcp", ":6632")
		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(e.Error())
			return
		}
	}()

	select {
	case <-time.After(35 * time.Second):
	case e := <-testResult:
		if e != nil {
			fmt.Println(e.Error())
			t.Fail()
		}
	}
}

func TestTcpX_TCP_AnchorMiddleware(t *testing.T) {
	var serverStart = make(chan int, 1)
	var testResult = make(chan error, 1)

	// client
	go func() {
		<-serverStart

		conn, err := net.Dial("tcp", "localhost:9004")
		if err != nil {
			testResult <- errorx.Wrap(err)
			fmt.Println(errorx.Wrap(err).Error())
			return
		}

		buf, e := PackJSON.Pack(1, "hello, I'm client")
		buf2, e := PackJSON.Pack(2, "hello, I'm client")

		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		conn.Write(buf)
		conn.Write(buf2)
	}()

	// server
	go func() {
		srv := NewTcpX(JsonMarshaller{})
		srv.OnMessage = nil
		srv.BeforeExit(func() {
			fmt.Println("exit")
		})
		// global middleware
		srv.UseGlobal(func(c *Context) {
			fmt.Println(1)
		})
		// anchor middleware
		srv.Use("anchor1", func(c *Context) {
			fmt.Println(2)
		})
		srv.AddHandler(2, func(c *Context) {
			fmt.Println(3)
		}, func(c *Context) {
			fmt.Println(4)
			c.Reply(10086, "hello, I'm server")
			testResult <- nil
		})
		srv.UnUse("anchor1")


		srv.Use("anchor1", func(c *Context) {
			fmt.Println(2)
		})

		srv.UnUse("anchor1")
		// router middleware
		srv.AddHandler(1, func(c *Context) {
			fmt.Println(3)
		}, func(c *Context) {
			fmt.Println(4)
			c.Reply(10086, "hello, I'm server")
			testResult <- nil
		})

		go func() {
			time.Sleep(time.Second * 10)
			serverStart <- 1
		}()
		e := srv.ListenAndServeTCP("tcp", ":9004")
		if e != nil {
			testResult <- errorx.Wrap(e)
			fmt.Println(e.Error())
			return
		}
	}()

	e := <-testResult
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
	}
}
