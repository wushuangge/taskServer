package _struct

//Service
type TaskService struct {
	ID				string 	`json:"id"`		     //Service id(唯一标识)
	Name			string 	`json:"name"`        //名称
	Type			string 	`json:"type"`        //类别
	URL				string 	`json:"url"`         //url
	Description		string 	`json:"description"` //描述
	Reserved		string 	`json:"reserved"`	 //预留
}

//HeartBeat
type HeartBeat struct {
	ID				string 	`json:"id"`		     //Service id(唯一标识)
	Reserved		string 	`json:"reserved"`	 //预留
}

//MetaData
type TaskMetaData struct {
	ID				string 	`json:"id"`			 //任务id(唯一标识)
	Name 			string 	`json:"name"`		 //任务名字
	URL				string 	`json:"url"`		 //url
	Description		string 	`json:"description"` //任务描述
	Status			string 	`json:"status"`		 //任务状态
	Reserved		string 	`json:"reserved"`	 //预留
}

//Manager
type TaskManager struct {
	ID		        string 	`json:"id"`			 //任务id(唯一标识)
	User	 		string 	`json:"user"`	 	 //用户
	GroupID 		string 	`json:"group_id"`	 //组
	Reserved		string 	`json:"reserved"`	 //预留
}