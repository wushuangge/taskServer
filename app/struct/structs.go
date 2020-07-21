package _struct

//Service
type TaskService struct {
	URL				string 	`bson:"_id"`         	//url
	Name			string 	`bson:"name"`        	//名称
	Reserved		string 	`bson:"reserved"`	 	//预留
}

//HeartBeat
type HeartBeat struct {
	URL				string 	`bson:"_id"`         	//url
	Time			int64 	`bson:"time"`		 	//发送心跳的时间
	Reserved		string 	`bson:"reserved"`	 	//预留
}

//Metadata
type TaskMetadata struct {
	TaskID      	string `bson:"_id"`     		//任务id(唯一标识)
	DataType        string `bson:"type"`   			//类别
	Status      	string `bson:"status"`      	//任务状态
	Reserved    	string `bson:"reserved"`    	//预留
}

//Manager
type TaskManagement struct {
	TaskID   		string `bson:"_id"`   			//任务id(唯一标识)
	DataType     	string `bson:"type"` 	  		//类别
	Status      	string `bson:"status"`   	 	//任务状态
	User     		string `bson:"user"`      		//用户
	Checker			string `bson:"checker"`   		//校验
	Group	  		string `bson:"group"`  			//组
	Reserved 		string `bson:"reserved"`  		//预留
}