package _struct

///任务
type TaskDash struct {
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

type TaskGroup struct {
	ID		        string 	`json:"id"`          //任务id(唯一标识)
	Name 			string 	`json:"name"`        //任务名称
	Addr 			string 	`json:"addr"`        //执行任务的微服务ip地址
	User	 		string 	`json:"user"`        //任务所属用户
	GroupID 		string 	`json:"group_id"`    //任务所属组
}