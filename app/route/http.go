package route

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"taskdash/app/store/mongodb"
	_struct "taskdash/app/struct"
)

func SetupHttp(g *gin.Engine)  {
	rget := g.Group("/rget")
	{
		rget.GET("/task", HandleTask)
		rget.GET("/test", HandleTest)
	}

	rpost := g.Group("/rpost")
	{
		rpost.POST("/task", HandleTask)
		rpost.POST("/test", HandleTest)
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
	case "GetTasksByUser":
		c.String(http.StatusOK, getTasksByUser("user", c.Request.FormValue("user")))
		break
	default:
		fmt.Println(c.Request.FormValue("operation"))
	}
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	taskMetadata := _struct.TaskMetadata{
		TaskID:      "sy-hn-2",
		DataType:    "2",
		Status:      "running",
		Reserved:    "no",
	}

	jsons, err := json.Marshal(taskMetadata)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, string(jsons))
}

func HttpPost(url string) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func HttpsPostForm(u string, operation string, createTime string) (string, error) {
	formData := url.Values{
		"operation":{operation},
		"createTime":{createTime},
	}
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm(u, formData)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
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
	response, err := mongodb.QueryConditionMetadata(key, value)
	if err != nil {
		return err.Error()
	}
	return response
}

func getTasksByUser(key string, value interface{}) string {
	response, err := mongodb.QueryConditionManagement(key, value)
	if err != nil {
		return err.Error()
	}
	return response
}

func getPagingTasks(limit int64, skip int64, key string, value interface{}) string {
	response, err := mongodb.QueryPagingMetadata(limit, skip, key, value)
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