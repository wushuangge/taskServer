package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/tcpx"
	"net"
	"os"
	"runtime/debug"
	"time"
)

const (
	request1	 		= "Register"
	request2			= "Update"
	MessageID 			= 2
)

///任务注册请求
type ReqTaskRegister struct {
	ID				string 	`json:"id"`
	Name 			string 	`json:"name"`
	URL				string 	`json:"url"`
	User	 		string 	`json:"user"`
	GroupID 		string 	`json:"group_id"`
	Description		string 	`json:"description"`
	Status			string 	`json:"status"`
	StartTime		string 	`json:"start_time"`
	StopTime		string 	`json:"stop_time"`
}

///任务注册响应
type RespTaskRegister struct {
	ID				string 	`json:"id"`
	Result 			string 	`json:"result"`
}

///任务状态更新请求
type ReqTaskStatus struct {
	ID				string 	`json:"id"`
	Status			string 	`json:"status"`
}

///任务状态更新响应
type RespTaskStatus struct {
	ID				string 	`json:"id"`
	Result 			string 	`json:"result"`
}


func main() {
	conn, e := net.Dial("tcp", "localhost:8090")
	if e != nil {
		panic(e)
	}
	go HeartBeat(conn)
	go Recv(conn)
	//go sendTask(conn)
	go sendStatus(conn)
	select {}
}

func sendTask(conn net.Conn){
	taskAttr := ReqTaskRegister{
		ID:				"1234567",
		Name: 			"abc",
		URL: 			"http://localhost:8080/v1/task",
		User: 			"wsg",
		GroupID: 		"mx",
		Description: 	"add data",
		Status: 		"running",
		StartTime: 		"12:00",
		StopTime:		"18:00",
	}

	jsons, err := json.Marshal(taskAttr)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("err : controllers - tool.go - StructI2json")
	}

	var msg []byte
	msg, err = tcpx.PackWithMarshallerName(tcpx.Message{
		MessageID: MessageID,
		Header: map[string]interface{}{
			"RequestID":		request1,
		},
		Body: string(jsons),
	}, "json")

	_, err = conn.Write(msg)
}

func sendStatus(conn net.Conn){
	taskStatus := ReqTaskStatus{
		ID:				"1234567",
		Status: 		"running",
	}

	jsons, err := json.Marshal(taskStatus)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("err : controllers - tool.go - StructI2json")
	}

	var msg []byte
	msg, err = tcpx.PackWithMarshallerName(tcpx.Message{
		MessageID: MessageID,
		Header: map[string]interface{}{
			"RequestID":		request2,
		},
		Body: string(jsons),
	}, "json")

	_, err = conn.Write(msg)
}

func Recv(conn net.Conn) {
	var buf []byte
	var e error
	for {
		buf, e = tcpx.FirstBlockOf(conn)
		if e != nil {
			handleError(e)
			return
		}
		messageID, e := tcpx.MessageIDOf(buf)
		if e != nil {
			handleError(e)
			return
		}
		body, e := tcpx.BodyBytesOf(buf)
		if e != nil {
			handleError(e)
			return
		}
		switch messageID {
		case 500, 400, 403:
			var m map[string]interface{}
			e := json.Unmarshal(body, &m)
			if e != nil {
				handleError(e)
				return
			}
			wellPrint(m)
		case 4:
			wellPrint(body)
		}
	}
}

func handleError(e error) {
	if e != nil {
		fmt.Printf("%v \n %s", e, debug.Stack())
		os.Exit(1)
	}
}

func wellPrint(src interface{}) {
	if buf, ok := src.([]byte); ok {
		fmt.Println(string(buf))
		return
	}

	buf, e := json.MarshalIndent(src, "  ", "  ")
	if e != nil {
		handleError(e)
	}
	fmt.Println(string(buf))
}

func HeartBeat(conn net.Conn) {
	var heartBeat []byte
	heartBeat, e := tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: tcpx.DEFAULT_HEARTBEAT_MESSAGEID,
		Header:    nil,
		Body:      nil,
	}, nil)
	for {
		_, e = conn.Write(heartBeat)
		if e != nil {
			fmt.Println(e.Error())
			break
		}
		time.Sleep(10 * time.Second)
	}
}
