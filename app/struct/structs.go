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
	ID      		string `bson:"_id"`     		//联合id(唯一标识)
	ProjectID  		string `bson:"project_id"`		//项目id
	InstanceID 		string `bson:"instance_id"`		//实例id
	TaskID     		string `bson:"task_id"`			//任务id
	TaskType        string `bson:"type"`   			//类别
	URL        		string `bson:"url"` 	  		//url
	Status      	string `bson:"status"`      	//任务状态
	CreateTime 		int64  `bson:"time"`      	    //创建时间
	Increment		int64  `bson:"increment"`      	//序列
	Reserved    	string `bson:"reserved"`    	//预留
}

//Manager
type TaskManagement struct {
	ID      		string `bson:"_id"`     		//联合id(唯一标识)
	ProjectID  		string `bson:"project_id"`		//项目id
	InstanceID 		string `bson:"instance_id"`		//实例id
	TaskID   		string `bson:"task_id"`   		//任务id(唯一标识)
	TaskType     	string `bson:"type"` 	  		//类别
	URL        		string `bson:"url"` 	  		//url
	Status      	string `bson:"status"`   	 	//任务状态
	CreateTime 		int64  `bson:"time"`      	    //创建时间
	User     		string `bson:"user"`      		//用户
	Checker			string `bson:"checker"`   		//校验
	Group	  		string `bson:"group"`  			//组
	Increment		int64  `bson:"increment"`      	//序列
	Reserved 		string `bson:"reserved"`  		//预留
}

type TaskFromService struct {
	ID         string `bson:"_id"`
	ProjectID  string
	InstanceID string
	TaskID     string
	TaskType   string
	Status     string
	URL        string
	ImgList    []string
	CreateTime int64
	Increment  int64
}
