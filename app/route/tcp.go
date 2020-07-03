package route

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/tcpx"
	"sync/atomic"
	"task/app/store"
	_struct "task/app/struct"
)

const (
	Register 				= "Register"
	Status 					= "Update"
	taskHandlerMessageID 	= 2
	clientMessageID 		= 4
)

var requestTimes int32

func countRequestTime(c *tcpx.Context) {
	atomic.AddInt32(&requestTimes, 1)
}

func getRequestTime(c *tcpx.Context) {
	c.JSON(clientMessageID, tcpx.H{"message": "success", "request_times": requestTimes})
}

func OnConnect(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("connecting from remote host %s network %s", c.ClientIP(), c.Network()))
}

func OnClose(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("disconnecting from remote host %s network %s", c.ClientIP(), c.Network()))
}

func OnHeartBeat(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("heartbeat from remote host %s network %s", c.ClientIP(), c.Network()))
	c.RecvHeartBeat()
}

func taskHandler(c *tcpx.Context) {
	var messageFromClient string
	var messageInfo tcpx.Message
	messageInfo, e := c.Bind(&messageFromClient)
	if e != nil {
		c.JSON(clientMessageID, "request format error")
		return
	}
	fmt.Println("header:", messageInfo.Header)
	fmt.Println("body:", messageInfo.Body)

	body := []byte(messageInfo.Body.(string))
	id := messageInfo.Header["RequestID"].(string)

	taskDash:=_struct.TaskDash{}
	err:=json.Unmarshal(body,&taskDash)
	if err!=nil{
		c.JSON(clientMessageID, "json.Unmarshal error")
		return
	}

	var res string
	switch (id){
	case Register:
		res = store.RegisterTask(taskDash.ID, body)
		break
	case Status:
		res = store.UpdateTask(taskDash.ID, body)
		break
	default:
		break
	}

	c.JSON(clientMessageID, res)
}
