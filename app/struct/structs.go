package _struct

//Service
type TaskService struct {
	URL				string 	`json:"url"`         //url
	Name			string 	`json:"name"`        //名称
	Type			string 	`json:"type"`        //类别
	Reserved		string 	`json:"reserved"`	 //预留
}

//HeartBeat
type HeartBeat struct {
	URL				string 	`json:"url"`         //url
	Time			int64 	`json:"time"`		 //发送心跳的时间
	Reserved		string 	`json:"reserved"`	 //预留
}

//Metadata
type TaskMetadata struct {
	ID		        string 	`json:"id"`          //任务id(唯一标识)
	Type			string 	`json:"type"`        //类别
	Description		string 	`json:"description"` //任务描述
	Status			string 	`json:"status"`      //任务状态
	Reserved		string 	`json:"reserved"`    //预留
}

//Manager
type TaskManagement struct {
	ID		        string 	`json:"id"`			 //任务id(唯一标识)
	User	 		string 	`json:"user"`	 	 //用户
	GroupID 		string 	`json:"group_id"`	 //组
	Reserved		string 	`json:"reserved"`	 //预留
}

//taskDash
type TaskDash struct {
	TaskID		    string 	`json:"task_id"`     		//任务id(唯一标识)
	TaskType		string 	`json:"task_type"`   		//任务类别
	TaskDescription	string 	`json:"task_description"` 	//任务描述
	TaskStatus		string 	`json:"task_status"`      	//任务状态
	ServiceName		string 	`json:"service_name"`		//服务名称
	ServiceURL		string 	`json:"service_url"` 		//服务url
	User	 		string 	`json:"user"`	 	 		//所属用户
	GroupID 		string 	`json:"group_id"`	 		//所属组
	Reserved		string 	`json:"reserved"`	 		//预留
}