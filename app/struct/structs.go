package _struct

//Service
type TaskService struct {
	ID       string            `bson:"_id"`      //id
	Network  string            `bson:"network"`  //网络类型
	Address  string            `bson:"address"`  //地址
	Name     string            `bson:"name"`     //名称
	Path     map[string]string `bson:"path"`     //路径
	Enable   bool              `bson:"enable"`   //是否启用
	Reserved string            `bson:"reserved"` //预留
}

//HeartBeat
type HeartBeat struct {
	ID       string `bson:"_id"`      //id
	Network  string `bson:"network"`  //网络类型
	Address  string `bson:"address"`  //地址
	Reserved string `bson:"reserved"` //预留
}

//Manager
type TaskManagement struct {
	ID          string `bson:"_id"`         //联合id(唯一标识)
	ProjectID   string `bson:"project_id"`  //项目id
	InstanceID  string `bson:"instance_id"` //实例id
	TaskID      string `bson:"task_id"`     //任务id
	TaskType    string `bson:"type"`        //类别
	URL         string `bson:"url"`         //url
	Status      string `bson:"status"`      //任务状态
	CreateTime  int64  `bson:"time"`        //创建时间
	User        string `bson:"user"`        //用户
	Distributor string `bson:"distributor"` //分配者
	Checker     string `bson:"checker"`     //校验
	Group       string `bson:"group"`       //组
	Reserved    string `bson:"reserved"`    //预留
}

type TaskFromService struct {
	ID         string `bson:"_id"`
	ProjectID  string
	InstanceID string
	TaskID     string
	TaskType   string
	Status     string
	URL        string
}

type ServiceTimer struct {
	Enable  bool
	Counter int
}

// Paging 分页结构
type Paging struct {
	Pos        int64         `json:"pos"`
	TotalCount int64         `json:"total_count"`
	Data       []interface{} `json:"data"`
}

type User struct {
	ID     string   `json:"_id"`    //用户代码
	Groups []string `json:"groups"` //组代码
	Level  int      `json:"level"`  //权限级别，0:editor，1:checker，2:manager
}

type Group struct {
	ID    string   `json:"_id"`   //组代码
	Users []string `json:"users"` //用户代码
}

const (
	LevelEditor  = 0x001
	LevelChecker = 0x010
	LevelManager = 0x100
)
