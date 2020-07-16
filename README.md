## TaskDash微服务

### 一.架构说明

![avatar](./source/1.png)

####1. 与任务服务交互：

任务服务通过push的方式将service信息push到nsq,服务管理模块通过pull的方式从nsq中取出服务信息，进行存储。

    type TaskService struct {
    	URL             string 	`json:"url"`            //url
    	Name            string 	`json:"name"`           //名称
    	Type            string 	`json:"type"`           //类别
    	Reserved        string 	`json:"reserved"`       //预留
    }
    
任务服务通过push的方式将心跳信息push到nsq,服务管理模块通过pull的方式从nsq中取出心跳信息进行校验。

    type HeartBeat struct {
    	URL             string 	`json:"url"`            //url
    	Time            int64 	`json:"time"`           //发送心跳的时间
    	Reserved        string 	`json:"reserved"`       //预留
    }

任务元数据模块通过service信息GET/POST到任务服务，获取任务元数据，然后就行存储。

    type TaskMetadata struct {
    	TaskID          string `json:"task_id"`         //任务id(唯一标识)
    	Type            string `json:"type"`            //类别
    	Description     string `json:"description"`     //任务描述
    	Status          string `json:"status"`          //任务状态
    	Reserved        string `json:"reserved"`        //预留
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

    type TaskManagement struct {
    	TaskID          string `json:"task_id"`         //任务id(唯一标识)
    	Type            string `json:"type"`            //类别
    	User            string `json:"user"`            //用户
    	GroupID         string `json:"group_id"`        //组
    	Reserved        string `json:"reserved"`        //预留
    }


### 二.接口说明

#####1.服务通知
    @title          service_push
    @network        nsq
    @topic:         taskService
    @Body:          json {"url":"http://localhost:8080/rpost/test","name":"abc","type":"2","reserved":"no"}
    
#####2.心跳通知
    @title          heartbeat_push
    @network        nsq 
    @topic:         heartBeat
    @Body:          json {"url":"http://localhost:8080/rpost/test","time":1594885152,"reserved":"no"}

#####3.查找请求
    @title          ReqFind
    @network        post
    @return         json {"task_id":"sy-hn-2","type":"2","description":"no","status":"running","reserved":"no"}
       
### 三.错误代码
    "errid":0          "errmsg":"success"
    "errid":1000       "errmsg":"请求格式不正确"
    "errid":1001       "errmsg":"数据库存储失败"
    "errid":1002       "errmsg":"数据库查询失败"