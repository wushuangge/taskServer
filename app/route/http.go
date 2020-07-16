package route

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"strings"
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
	case "userLogin":
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")
		userIsExist := checkUser(username, password)
		if userIsExist == true {
			c.String(http.StatusOK, "success")
		} else {
			c.String(http.StatusOK, "failure")
		}
		break
	default:
		fmt.Println("default")
	}
}

func HandleTest(c *gin.Context)  {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	taskMetadata := _struct.TaskMetadata{
		TaskID:      "sy-hn-2",
		Type:        "2",
		Description: "no",
		Status:      "running",
		Reserved:    "no",
	}

	jsons, err := json.Marshal(taskMetadata)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, string(jsons))
}

func HttpPost(url string) (string, error){
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

func HttpPostWithCookie(url string, cookie string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
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