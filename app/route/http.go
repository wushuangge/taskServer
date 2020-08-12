package route

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func SetupHttp(g *gin.Engine)  {
	v1 := g.Group("/rget")
	{
		v1.GET("/task", HandleTask)
		v1.GET("/test", HandleTest)
	}

	v2 := g.Group("/rpost")
	{
		v2.POST("/task", HandleTask)
		v2.POST("/test", HandleTest)
	}
}

func HandleTask(c *gin.Context)  {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	switch c.Request.FormValue("operation") {
	case "UserLogin":
		c.String(http.StatusOK, userLogin(c.Request.FormValue("username"),
			c.Request.FormValue("password")))
		break
	case "GetTasksByStatus":
		c.String(http.StatusOK, getTasksByStatus("status", c.Request.FormValue("status")))
		break
	case "GetTasksNoUser":
		c.String(http.StatusOK, getTasksNoUser())
		break
	case "GetTasksByUser":
		c.String(http.StatusOK, getTasksByUser("user", c.Request.FormValue("username")))
		break
	case "TaskBind":
		c.String(http.StatusOK,taskBind(c.Request.FormValue("id"), c.Request.FormValue("username")))
		break
	default:
		log.Error("default!!! operation is ",c.Request.FormValue("operation"))
	}
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	taskMetadata := _struct.TaskMetadata{
		TaskID:      "sy-hn-2",
		TaskType:    "2",
		Status:      "running",
		Reserved:    "no",
	}

	jsons, err := json.Marshal(taskMetadata)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, string(jsons))
}

func userLogin(username string, password string) string {
	userIsExist := checkUser(username, password)
	var result string
	if userIsExist == true {
		result = "success"
	} else {
		result = "error"
	}
	return result
}

func getTasksByStatus(key string, value interface{}) string {
	response, err := mongodb.QueryConditionManagement(bson.D{{key, value}})
	if err != nil {
		return err.Error()
	}
	return response
}

func getTasksNoUser() string {
	response, err := mongodb.QueryConditionManagement(bson.D{{"user", bson.M{"$exists": false}}})
	if err != nil {
		return err.Error()
	}
	return response
}

func getTasksByUser(key string, value interface{}) string {
	response, err := mongodb.QueryConditionManagement(bson.D{{key, value}})
	if err != nil {
		return err.Error()
	}
	return response
}

func taskBind(id string, user string) string {
	filter :=bson.M{"_id": id}

	update := bson.D{
		{"$set", bson.D{
			{"status", "已领取"},
			{"user", user},
		}},
	}
	err := mongodb.UpdateManagement(filter, update, false)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func getPagingTasks(limit int64, skip int64, key string, value interface{}) string {
	response, err := mongodb.QueryPagingManagement(limit, skip, bson.D{{key, value}})
	if err != nil {
		return err.Error()
	}
	return response
}

func checkUser(username string, password string) bool {
	var userIsExist = false
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, _, errs := request.Get("https://192.168.51.30:5001/auth").SetBasicAuth(username, password).Set("User-Agent", "ftp").End()
	if len(errs) <= 0 && resp.StatusCode == 200 {
		userIsExist = true
	}
	return userIsExist
}
