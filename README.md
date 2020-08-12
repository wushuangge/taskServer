## TaskDash微服务

### 一.架构说明

![avatar](source/1.png)

####1. 与任务服务交互：

任务服务通过push的方式将service信息push到nsq,服务管理模块通过pull的方式从nsq中取出服务信息，进行存储。

    type TaskService struct {
    	URL             string 	`bson:"_id"`            //url
    	Name            string 	`bson:"name"`           //名称
    	Reserved        string 	`bson:"reserved"`       //预留
    }
    
任务服务通过push的方式将新任务push到nsq,服务管理模块通过pull的方式从nsq中取出新任务，进行存储及其他处理。

    type TaskFromService struct {
    	ID              string `bson:"_id"`
    	ProjectID       string
    	InstanceID      string
    	TaskID          string
    	Status          string
    	TaskType        string
    	URL             string
    	EditInfo        string
    	CreateTime      int64
    }
    
任务服务通过push的方式将任务更新信息push到nsq,服务管理模块通过pull的方式从nsq中取出任务更新信息，进行任务更新。

    type TaskFromService struct {
    	ID              string `bson:"_id"`
    	ProjectID       string
    	InstanceID      string
    	TaskID          string
    	Status          string
    	TaskType        string
    	URL             string
    	EditInfo        string
    	CreateTime      int64
    }
    
任务服务通过push的方式将心跳信息push到nsq,服务管理模块通过pull的方式从nsq中取出心跳信息进行校验。

    type HeartBeat struct {
    	URL             string 	`bson:"_id"`            //url
    	Time            int64 	`bson:"time"`           //发送心跳的时间
    	Reserved        string 	`bson:"reserved"`       //预留
    }

任务元数据模块通过service信息GET/POST到任务服务，获取任务元数据，然后就行存储。

    type TaskMetadata struct {
    	ID              string `bson:"_id"`             //联合id(唯一标识)
    	ProjectID       string `bson:"project_id"`      //项目id
    	InstanceID      string `bson:"instance_id"`     //实例id
    	TaskID          string `bson:"task_id"`         //任务id
    	TaskType        string `bson:"type"`            //类别
    	URL             string `bson:"url"`             //url
    	Status          string `bson:"status"`          //任务状态
    	CreateTime      int64  `bson:"time"`            //创建时间
    	Reserved        string `bson:"reserved"`        //预留
    }
      
####2. 与任务管理前端交互：

任务查询

web前端先进行登录验证，然后通过GET/POST的方式从任务管理服务查询任务信息并显示。

示例：
User：smq 
前端查询任务返回如下：

| ID   | ProjectID | InstanceID | TaskID |TaskType|URL  | Status |CreateTime
|------|-----------|------------|--------|--------|-----|--------|----------|
| 133  | i1        | p1         | 0001   |satEdit |?    |编辑中   |1596174701|
| 132  | i1        | p2         | 0002   |satEdit |?    |编辑中   |1596174701|
| 131  | i1        | p3         | 0003   |satEdit |?    |编辑中   |1596174701|
| 130  | i1        | p4         | 0004   |satEdit |?    |编辑中   |1596174701|
| 137  | i1        | p5         | 0005   |satEdit |?    |编辑中   |1596174701|

用户进行任务领取，web前端向任务管理服务发送任务分配信息，任务管理服务将用户ID和任务ID进行绑定然后存储。

    type TaskManagement struct {
    	ID              string `bson:"_id"`             //联合id(唯一标识)
    	ProjectID       string `bson:"project_id"`      //项目id
    	InstanceID      string `bson:"instance_id"`     //实例id
    	TaskID          string `bson:"task_id"`         //任务id
    	TaskType        string `bson:"type"`            //类别
    	URL             string `bson:"url"`             //url
    	Status          string `bson:"status"`          //任务状态
    	CreateTime      int64  `bson:"time"`            //创建时间
    	User            string `bson:"user"`            //用户
    	Checker         string `bson:"checker"`         //校验
    	Group           string `bson:"group"`           //组
    	Reserved        string `bson:"reserved"`        //预留
    }
    
mongodb中json存储：
    
    {
    	"_id" : "24079070c4f1a4bc9b57070ae831514f",
    	"instance_id" : "instanceID",
    	"project_id" : "projectID",
    	"status" : "待领取",
    	"task_id" : "20200311_0000001_111111",
    	"time" : 1596174701,
    	"type" : "vidEdit",
    	"url" : "https://192.168.51.33:8080",
    	"user" : "sunmq"
    }
    
高级用户进行任务指派:TODO

### 二.服务发现接口说明

#####1.服务注册
    @title          service_push
    @network        nsq
    @topic:         serviceRegister
    @Body:          json {"URL":"https://192.168.51.33:8080/TaskDash/","Name":"annotator","Reserved":""}
    
#####1.任务注册
    @title          task_push
    @network        nsq
    @topic:         taskRegister
    @Body:          json {"ID":"d8b6efc3ee308ef3ea1284b9dfc79df8","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_020","TaskType":"satEdit","Status":"待领取","URL":"https://192.168.51.33:8080","EditInfo","CreateTime":1597116419}
    
#####1.任务更新
    @title          update_push
    @network        nsq
    @topic:         updateStatus
    @Body:          json {"ID":"d8b6efc3ee308ef3ea1284b9dfc79df8","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_020","TaskType":"satEdit","Status":"编辑中","URL":"https://192.168.51.33:8080","EditInfo","CreateTime":1597116419}
           
#####2.心跳通知
    @title          heartbeat_push
    @network        nsq 
    @topic:         heartBeat
    @Body:          json {"url":"http://localhost:8080/rpost/test","time":1594885152,"reserved":"no"}

       
### 三.前端接口说明

#####1.1登录

    @title          UserLogin
    @network        post
    @parameter1     operation   :   "userLogin"
    @parameter2     username    :   "abc"
    @parameter2     password    :   "123456"
    @return         string

#####1.2请求示例

    var frameData = {
        "operation":"userLogin",
        "username":VarTool.GetCookie("username"),
        "password":VarTool.GetCookie("password")
    };
            
    var resData = AjaxTool.SendData(frameData,"/rpost/task");
    
#####1.3返回示例

    "success"
    
#####2.1获取未领取的任务列表

    @title          GetTasksNoUser
    @network        post
    @parameter1     operation   :   "GetTasksNoUser"
    @return         string
    
#####2.2请求示例

    var frameData = {
        "operation":"GetTasksNoUser"
    };
            
    var resData = AjaxTool.SendData(frameData,"/rpost/task");
    
#####2.3返回示例

    [{"ID":"9cf4db018beecb08a1f27678798c0653","ProjectID":"p1","InstanceID":"i1","TaskID":"20200311_0000001_111111","TaskType":"vidEdit","URL":"https://192.168.51.33:8080","Status":"待领取","CreateTime":1596792596,"User":"","Checker":"","Group":"","Reserved":""}]
    
#####3.1获取对应用户下任务列表

    @title          GetUserTasks
    @network        post
    @parameter1     operation   : "GetTasksByUser"
    @parameter2     user        : "smq"
    @return         string
    
#####3.2请求示例

    var frameData = {
        "operation":"GetTasksByUser",
        "user":"smq"
    };
            
    var resData = AjaxTool.SendData(frameData,"/rpost/task");
    
#####3.3返回示例

    [{"ID":"777417adfc7aaa032537e75b025ab267","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_022","TaskType":"satEdit","URL":"https://192.168.51.33:8080","Status":"编辑中","CreateTime":1595827431,"User":"sunmq","Checker":"","Group":"","Reserved":""},
    {"ID":"7c1c5ed25bf0c827bd8ddcb1091fad1c","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_0222","TaskType":"satEdit","URL":"https://192.168.51.33:8080","Status":"编辑中","CreateTime":1595813604,"User":"sunmq","Checker":"","Group":"","Reserved":""},
    {"ID":"24079070c4f1a4bc9b57070ae831514f","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"20200311_0000001_111111","TaskType":"vidEdit","URL":"https://192.168.51.33:8080","Status":"待领取","CreateTime":1596174701,"User":"sunmq","Checker":"","Group":"","Reserved":""},
    {"ID":"fdbcc3906f2c619394b94d636af353a5","ProjectID":"p1","InstanceID":"i1","TaskID":"0210232_1220_1222_0222","TaskType":"satEdit","URL":"https://192.168.51.33:8080","Status":"已领取","CreateTime":1596764687,"User":"sunmq","Checker":"","Group":"","Reserved":""}]

#####4.1领取任务

    @title          TaskBind
    @network        post
    @parameter1     operation   : "TaskBind"
    @parameter2     user        : "smq"
    @return         string
    
#####4.2请求示例

    var frameData = {
        "operation":"TaskBind",
        "user":"smq"
    };
            
    var resData = AjaxTool.SendData(frameData,"/rpost/task");
    
#####4.3返回示例

    "success"
  
#####5.1获取对应状态下的任务列表

    @title          GetTasksByStatus
    @network        post
    @parameter1     operation   :   "GetTasksByStatus"
    @parameter2     status      :   "编辑中"
    @return         string
    
#####5.2请求示例

    var frameData = {
        "operation": "GetTasksByStatus",
        "status"   : "编辑中"
    };
            
    var resData = AjaxTool.SendData(frameData,"/rpost/task");
    
#####5.3返回示例

        [{"ID":"777417adfc7aaa032537e75b025ab267","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_022","TaskType":"satEdit","URL":"https://192.168.51.33:8080","Status":"编辑中","CreateTime":1595827431,"User":"sunmq","Checker":"","Group":"","Reserved":""},
        {"ID":"7c1c5ed25bf0c827bd8ddcb1091fad1c","ProjectID":"projectID","InstanceID":"instanceID","TaskID":"0210232_1220_1222_0222","TaskType":"satEdit","URL":"https://192.168.51.33:8080","Status":"编辑中","CreateTime":1595813604,"User":"sunmq","Checker":"","Group":"","Reserved":""}]

    
### 四.错误代码

    "errid":0          "errmsg":"success"
    "errid":1000       "errmsg":"请求格式不正确"
    "errid":1001       "errmsg":"数据库存储失败"
    "errid":1002       "errmsg":"数据库查询失败"