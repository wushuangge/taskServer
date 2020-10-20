package route

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"strconv"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func SetupHttp(r *gin.Engine) {
	v1 := r.Group("/rget")
	{
		v1.GET("/task", HandleTask)
		v1.GET("/test", HandleTest)
	}

	v2 := r.Group("/rpost")
	{
		v2.POST("/task", HandleTask)
		v2.POST("/test", HandleTest)
	}

	v3 := r.Group("/ui")
	{
		v3.StaticFS("/", http.Dir("./ui"))
	}
}

func HandleTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	auth, user := authUser(c)
	if !auth {
		c.String(http.StatusOK, "该用户未登录")
		return
	}
	switch c.Request.FormValue("operation") {
	case "UserLogin":
		result, attrs, code := userLogin(c.Request.FormValue("username"), c.Request.FormValue("password"))
		if code == http.StatusOK {
			session := sessions.Default(c)
			session.Set("user", _struct.User{c.Request.FormValue("username"), attrs, getLevel(attrs)})
			session.Save()
		}
		c.String(code, result)
		break
	case "UserLogout":
		session := sessions.Default(c)
		session.Delete("user")
		session.Save()
		c.String(http.StatusOK, "success")
		break
	case "GetTasksNoUser":
		start, err1 := strconv.ParseInt(c.Request.FormValue("start"), 10, 64)
		count, err2 := strconv.ParseInt(c.Request.FormValue("count"), 10, 64)
		if err1 != nil || err2 != nil {
			start = 0
			count = 50
		}
		c.String(http.StatusOK, getPagingTasksNoUser(count, start, user.Level))
		break
	case "GetTasksByUser":
		start, err1 := strconv.ParseInt(c.Request.FormValue("start"), 10, 64)
		count, err2 := strconv.ParseInt(c.Request.FormValue("count"), 10, 64)
		if err1 != nil || err2 != nil {
			start = 0
			count = 50
		}

		c.String(http.StatusOK, getPagingTasksByUser(count, start, "user", user.ID, user.Level))
		break
	case "GetCompleteTasksByUser":
		start, err1 := strconv.ParseInt(c.Request.FormValue("start"), 10, 64)
		count, err2 := strconv.ParseInt(c.Request.FormValue("count"), 10, 64)
		if err1 != nil || err2 != nil {
			start = 0
			count = 50
		}
		c.String(http.StatusOK, getPagingCompleteTasksByUser(count, start, "user", user.ID, user.Level))
		break
	case "TaskBind":
		c.String(http.StatusOK, taskBind(c.Request.FormValue("id"), user.ID))
		break
	case "GetServiceUrl":
		c.String(http.StatusOK, getServiceUrl(c.Request.FormValue("name"), c.Request.FormValue("path")))
		break
	case "GetUsers":
		c.String(http.StatusOK, getUsers(c.Request.FormValue("group")))
		break
	case "GetGroups":
		c.String(http.StatusOK, getGroups(c.Request.FormValue("user")))
		break
	case "GetEditors":
		c.String(http.StatusOK, getEditors(c.Request.FormValue("user")))
		break
	case "GetGroupMembers":
		c.String(http.StatusOK, getGroupMembers(c.Request.FormValue("group")))
		break
	default:
		log.Error("default!!! operation is ", c.Request.FormValue("operation"))
	}
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, string("hello world!"))
}

func userLogin(username string, password string) (string, []string, int) {
	attrs, err := getUserAttrs(username, password)
	if err != nil {
		return err.Error(), []string{}, http.StatusBadRequest
	}
	result, err := json.Marshal(attrs["groups"])
	if err != nil {
		return err.Error(), []string{}, http.StatusBadRequest
	}
	return string(result), attrs["groups"], http.StatusOK
}

func getPagingTasksNoUser(limit int64, skip int64, level int) string {
	taskType := getTaskType(level)
	filter := bson.D{{"user", bson.M{"$exists": false}},
		{"type", primitive.Regex{Pattern: taskType}}}
	num, err := mongodb.QueryManagementNum(filter)
	if err != nil {
		return err.Error()
	}
	response, err := mongodb.QueryPagingManagement(limit, skip, filter)
	if err != nil {
		return err.Error()
	}
	paging := _struct.Paging{
		Pos:        skip,
		TotalCount: num,
		Data:       response,
	}
	encode, err := json.Marshal(paging)
	if err != nil {
		return err.Error()
	}
	return string(encode)
}

func getPagingTasksByUser(limit int64, skip int64, key string, value interface{}, level int) string {
	taskType := getTaskType(level)
	filter := bson.D{{key, value},
		{"type", primitive.Regex{Pattern: taskType}}}
	num, err := mongodb.QueryManagementNum(filter)
	if err != nil {
		return err.Error()
	}
	response, err := mongodb.QueryPagingManagement(limit, skip, filter)
	if err != nil {
		return err.Error()
	}
	paging := _struct.Paging{
		Pos:        skip,
		TotalCount: num,
		Data:       response,
	}
	encode, err := json.Marshal(paging)
	if err != nil {
		return err.Error()
	}
	return string(encode)
}

func getPagingCompleteTasksByUser(limit int64, skip int64, key string, value interface{}, level int) string {
	taskType := getTaskType(level)
	filter := bson.D{{key, value},
		{"status", "已完成"},
		{"type", primitive.Regex{Pattern: taskType}}}
	num, err := mongodb.QueryBackupNum(filter)
	if err != nil {
		return err.Error()
	}
	response, err := mongodb.QueryPagingBackup(limit, skip, filter)
	if err != nil {
		return err.Error()
	}
	paging := _struct.Paging{
		Pos:        skip,
		TotalCount: num,
		Data:       response,
	}
	encode, err := json.Marshal(paging)
	if err != nil {
		return err.Error()
	}
	return string(encode)
}

func taskBind(id string, user string) string {
	res, err := mongodb.QueryConditionManagement(bson.D{{"_id", id},
		{"user", user}})
	if res != "[]" && err == nil {
		return "this id is already in use," + res
	}
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", "已领取"},
			{"user", user},
		}},
	}
	err = mongodb.UpdateManagement(filter, update, false)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func taskAssigned(id string, user string, distributor string) string {
	res, err := mongodb.QueryConditionManagement(bson.D{{"_id", id},
		{"user", user}})
	if res != "[]" && err == nil {
		return "this id is already in use," + res
	}
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"distributor", distributor},
			{"status", "已指派"},
			{"user", user},
		}},
	}
	err = mongodb.UpdateManagement(filter, update, false)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func getUsers(group string) string {
	filter := bson.D{{"_id", group}}
	response, err := mongodb.QueryConditionGroup2json(filter)
	if err != nil {
		return err.Error()
	}
	return response
}

func getGroups(user string) string {
	filter := bson.D{{"_id", user}}
	response, err := mongodb.QueryConditionUser2json(filter)
	if err != nil {
		return err.Error()
	}
	return response
}

func getEditors(user string) string {
	filter := bson.D{}
	response, err := mongodb.QueryConditionGroup2Struct(filter)
	if err != nil {
		return err.Error()
	}

	jsons, err := json.Marshal(response[0].Users)
	if err != nil {
		return ""
	}
	return string(jsons)
}

func getGroupMembers(group string) string {
	response := getGroupAttrs(group)
	jsons, err := json.Marshal(response)
	if err != nil {
		return ""
	}
	return string(jsons)
}

func getServiceUrl(name string, path string) string {
	filter := bson.D{{"name", name},
		{"enable", true}}
	response, err := mongodb.QueryConditionService2Struct(filter)
	if err != nil {
		return err.Error()
	}
	if len(response[0].Network) == 0 || len(response[0].Address) == 0 ||
		len(response[0].Path[path]) == 0 {
		return "error"
	}

	url := response[0].Network + response[0].Address + "/" +
		response[0].Path[path] + "/"
	fmt.Println("getServiceUrl:", url)
	return url
}

func authUser(c *gin.Context) (bool, _struct.User) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		if c.Request.FormValue("operation") == "UserLogin" {
			return true, _struct.User{}
		} else {
			return false, _struct.User{}
		}
	}
	return true, user.(_struct.User)
}

func checkUser(username string, password string) bool {
	var userIsExist = false
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, _, errs := request.Get("https://192.168.51.33:5001/auth").SetBasicAuth(username, password).Set("User-Agent", "ftp").End()
	if len(errs) <= 0 && resp.StatusCode == 200 {
		userIsExist = true
	}
	return userIsExist
}

// getUserAttrs 获取用户属性
func getUserAttrs(username string, password string) (map[string][]string, error) {
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, body, errs := request.Get("https://192.168.51.33:5010/auth").SetBasicAuth(username, password).Set("User-Agent", "ftp").End()
	if len(errs) > 0 || resp.StatusCode != 200 {
		return nil, errors.New("get user err")
	}
	var attrs = make(map[string][]string)
	json.Unmarshal([]byte(body), &attrs)
	filter := bson.M{"_id": username}
	update := bson.D{
		{"$set", bson.D{
			{"_id", username},
			{"Groups", attrs["groups"]},
		}},
	}
	mongodb.UpdateUser(filter, update, true)
	return attrs, nil
}

func getGroupAttrs(group string) [][]string {
	var result = make([][]string, 0)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	url := "https://192.168.51.33:5010/users?cn=" + group
	resp, err := client.Get(url)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var entry = make([]map[string][]string, 0)
	err = json.Unmarshal(body, &entry)
	if err != nil {
		return result
	}
	for _, v := range entry {
		result = append(result, v["uid"])
	}
	return result
}

func getGroupsInfo(group string) {
	entry := getGroupAttrs(group)
	for _, users := range entry {
		for _, value := range users {
			filter := bson.M{"_id": value}
			update := bson.D{
				{"$set", bson.D{
					{"_id", value},
					{"Groups", group},
				}},
			}
			mongodb.UpdateUser(filter, update, true)
		}
	}
}

func getLevel(groups []string) int {
	level := 0x000
	for _, v := range groups {
		switch v {
		case "editor":
			level |= _struct.LevelEditor
			break
		case "checker":
			level |= _struct.LevelChecker
			break
		case "manager":
			level |= _struct.LevelManager
			break
		}
	}
	return level
}

func getTaskType(level int) string {
	if level&0x010 == _struct.LevelChecker {
		return "Check"
	} else if level&0x001 == _struct.LevelEditor {
		return "Edit"
	} else if level&0x100 == _struct.LevelManager {
		return ""
	}
	return ""
}

func LoadGroupsInfo() {
	getGroupsInfo("editor")
	getGroupsInfo("manager")
	getGroupsInfo("checker")
}
