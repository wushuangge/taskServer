### TaskDash微服务

####架构说明
TaskDash对外提供Register接口

### TaskDash微服务接口说明

####结构定义
    type TaskAttr struct {
	ID	            string
	Name            string
	URL             string
	UserName 		string
	GroupID 		string
	Description		string
	Status			string
	StartTime		string
	StopTime		string
    }


####注册
    @title       Register
    @description 通过rpc的方式调用注册接口
    @param       input TaskAttr
    @param       output TaskAttr
    @return      error
    func (server *Server) Register(input *TaskAttr, output *TaskAttr) error
    
####状态通知
    @title       OnStatus
    @description 通过rpc的方式调用状态通知
    @param       input TaskAttr
    @return      error
    func (server *Server) OnStatus(input _struct.TaskAttr) error

####心跳
    @title       OnHeartbeat
    @description 通过rpc的方式调用状态通知
    @param       input TaskAttr
    @return      error
    func (server *Server) OnHeartbeat(input _struct.TaskAttr) error 

