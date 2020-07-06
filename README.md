## TaskDash微服务

### 一.架构说明

####1. 与其他微服务交互：

TaskDash接收其他微服务任务注册请求，存储微服务任务信息，请求方式为tcp请求，详细请求结构参考如下：

#####任务属性结构体定义：
    type TaskDash struct {
    	ID		        string 	`json:"id"`          //任务id(唯一标识)
    	Name 			string 	`json:"name"`        //任务名称
    	URL		        string 	`json:"url"`         //任务url
    	User	 		string 	`json:"user"`        //任务所属用户
    	GroupID 		string 	`json:"group_id"`    //任务所属组
    	Description		string 	`json:"description"` //任务描述
    	Status			string 	`json:"status"`      //任务状态
    	StartTime		string 	`json:"start_time"`  //任务开始时间
    	StopTime		string 	`json:"stop_time"`   //任务结束时间
    }
    
建立tcp连接：连接TaskDash服务 -> 连接成功 -> 发送心跳 -> end
    
注册：请求注册 -> TaskDash处理注册 -> 响应注册结果 -> end

更新：请求更新 -> TaskDash处理更新 -> 响应更新结果 -> end

查询：请求更新 -> TaskDash处理更新 -> 响应更新结果 -> end

心跳：若长时间未收到心跳请求，TaskDash认为当前的连接已经终止，更新任务状态为shutdown，待重新连接后查询状态更新数据库

具体请求方式详见下方接口说明
    
####2. 与任务管理前端交互：

TaskDash也为前端提供任务管理信息显示，根据前端用户的权限（用户id，组id）来提供相关的任务信息列表。

登录：前端用户登录 -> TaskDash请求认证授权 -> TaskDash根据认证授权结果返回登录成功or失败 -> end

查询：前端用户请求查询任务 -> TaskDash判断前端用户是否已登录 -> TaskDash判断用户所属组 ->TaskDash查询数据库返回查询结果 -> end

示例：
User：ssx 
GroupID：mx 
前端查询任务返回如下：

| ID   | Name | URL | User | GroupID | Description |Status  |StartTime  |StopTime   |
|------|------|-----|------|---------|-------------|--------|-----------|-----------|
| 133  | zxf  | ?   | ssx  | mx      | ?           |running |2020.07.01 |2020.07.05 |
| 132  | lxq  | ?   | ssx  | mx      | ?           |finish  |2020.07.16 |2020.07.27 |            
| 131  | lb   | ?   | ssx  | mx      | ?           |shutdown|2020.07.03 |2020.07.06 |
| 130  | zgl  | ?   | ssx  | mx      | ?           |restart |2020.07.04 |2020.07.15 |
| 137  | zy   | ?   | ssx  | mx      | ?           |running |2020.07.15 |2020.07.18 |

同时也支持其他前端扩展，待定。

####3. 内部逻辑处理
#####用户组管理结构定义
    type TaskGroup struct {
    	ID		        string 	`json:"id"`          //任务id(唯一标识)
    	Name 			string 	`json:"name"`        //任务名称
    	Addr 			string 	`json:"addr"`        //执行任务的微服务ip地址
    	User	 		string 	`json:"user"`        //任务所属用户
    	GroupID 		string 	`json:"group_id"`    //任务所属组
    }

### 二.接口说明

#####1.注册请求
    @title       ReqRegister
    @network     tcp请求
    @package        
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Register"}
        Body:       json {"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"}

#####2.注册响应
    @title       RespRegister
    @network     tcp请求
    @package    
        body        json {"id":"1234567","errid":0,"errmsg":"正确"}
    
#####3.更新请求
    @title       ReqUpdate
    @network     tcp请求
    @package        
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Update"}
        Body:       json {"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"}

#####4.更新响应
    @title       RespUpdate
    @network     tcp请求
    @package    
        body        json {"id":"1234567","errid":0,"errmsg":"正确"}
    
#####5.查找请求
    @title       ReqFind
    @network     tcp请求
    @package        
        MessageID:  int 2
        Header:     map[string]interface{} {"RequestID" : "Find"}
        Body:       json {"id":"1234567"}

#####6.查找响应
    @title       RespFind
    @network     tcp请求
    @package    
        body        json {{"id":"1234567","name":"abc","url":"http://localhost:8080/v1/task","user":"wsg","group_id":"mx","description":"add data","status":"running","start_time":"12:00","stop_time":"18:00"},"errid":0,"errmsg":"正确"}
    

#####7.心跳通知
    @title       Heartbeat
    @network     tcp请求
    @package        
        MessageID:  int 1392
        Header:     nil
        Body:       nil
        
### 三.错误代码
    "errid":0          "errmsg":"正确"
    "errid":1000       "errmsg":"请求格式不正确"
    "errid":1001       "errmsg":"数据库存储失败"
    "errid":1002       "errmsg":"不存在该任务id"