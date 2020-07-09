## TaskDash微服务

### 一.架构说明

![avatar](./source/1.png)

####1. 与任务服务交互：

任务服务通过push的方式将service信息push到nsq,服务管理模块通过pull的方式从nsq中取出服务信息，进行存储。

    type TaskService struct {
    	ID              string 	`json:"id"`          //Service id(唯一标识)
    	Name            string 	`json:"name"`        //名称
    	Type            string 	`json:"type"`        //类别
    	URL             string 	`json:"url"`         //url
    	Description     string 	`json:"description"` //描述
    	Reserved        string 	`json:"reserved"`    //预留
    }
    
任务服务通过push的方式将心跳信息push到nsq,服务管理模块通过pull的方式从nsq中取出心跳信息进行校验。

    type HeartBeat struct {
    	ID		        string 	`json:"id"`          //Service id(唯一标识)
    	Reserved		string 	`json:"reserved"`    //预留
    }

任务元数据模块通过service信息GET/POST到任务服务，获取任务元数据，然后就行存储。

    type TaskMetaData struct {
    	ID		        string 	`json:"id"`          //任务id(唯一标识)
    	Name 			string 	`json:"name"`        //任务名字
    	URL 			string 	`json:"url"`         //url
    	Description		string 	`json:"description"` //任务描述
    	Status			string 	`json:"status"`      //任务状态
    	Reserved		string 	`json:"reserved"`    //预留
    }
    
####2. 与任务管理前端交互：

任务查询

web前端先进行登录验证，然后通过GET/POST的方式从任务管理服务查询任务信息并显示。

示例：
User：ssx 
前端查询任务返回如下：

| ID   | Name | URL | Description |Status  |Reserved  |
|------|------|-----|-------------|--------|----------|
| 133  | zxf  | ?   | ?           |running |?         |
| 132  | lxq  | ?   | ?           |finish  |?         |
| 131  | lb   | ?   | ?           |shutdown|?         |
| 130  | zgl  | ?   | ?           |restart |?         |
| 137  | zy   | ?   | ?           |running |?         |

用户进行任务领取（权限更高的用户可以进行任务分配），web前端向任务管理服务发送任务分配信息，
任务管理服务将用户ID和任务ID进行绑定然后存储。

    type TaskManager struct {
    	ID		        string 	`json:"id"`          //任务id(唯一标识)
    	User	 		string 	`json:"user"`        //用户
    	GroupID 		string 	`json:"group_id"`    //组(预留)
    	Reserved		string 	`json:"reserved"`    //预留
    }

### 二.接口说明

#####1.注册请求
    @title       ReqRegister
    @network     tcp请求
    @pack
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Register"}
        Body:       json {"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"}

#####2.注册响应
    @title       RespRegister
    @network     tcp请求
    @pack
        body        json {"id":"1234567","errid":0,"errmsg":"正确"}
    
#####3.更新请求
    @title       ReqUpdate
    @network     tcp请求
    @pack     
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Update"}
        Body:       json {"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"}

#####4.更新响应
    @title       RespUpdate
    @network     tcp请求
    @pack
        body        json {"id":"1234567","errid":0,"errmsg":"正确"}
    
#####5.查找请求
    @title       ReqFind
    @network     tcp请求
    @pack    
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Find"}
        Body:       json {"id":"1234567"}

#####6.查找响应
    @title       RespFind
    @network     tcp请求
    @pack
        body        json {{"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"},"errid":0,"errmsg":"正确"}
    

#####7.心跳通知
    @title       Heartbeat
    @network     tcp请求
    @pack   
        MessageID:  int 1392
        Header:     nil
        Body:       nil
        
### 三.错误代码
    "errid":0          "errmsg":"正确"
    "errid":1000       "errmsg":"请求格式不正确"
    "errid":1001       "errmsg":"数据库存储失败"
    "errid":1002       "errmsg":"不存在该任务id"