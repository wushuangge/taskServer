package route

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"task/app/store/mongodb"
	_struct "task/app/struct"
)

func SetupHttp(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		v1.GET("/task", HandleGetMetadata)
		v1.GET("/test", HandleGetTest)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/task", HandlePostMetadata)
		v2.POST("/test", HandleGetTest)
	}
}

func HandleGetMetadata(c *gin.Context)  {
	response, err :=mongodb.QueryAllMetadata()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, response)
}

func HandlePostMetadata(c *gin.Context)  {
	response, err :=mongodb.QueryAllMetadata()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, response)
}

func HandleGetTest(c *gin.Context)  {
	taskMetadata := _struct.TaskMetadata{
		ID: 			"sy-hn-1",
		Type:			"1",
		Description: 	"no",
		Status:			"running",
		Reserved:		"no",
	}

	jsons, err := json.Marshal(taskMetadata)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, string(jsons))
}

func HandlePostTest(c *gin.Context)  {
	taskMetadata := _struct.TaskMetadata{
		ID: 			"sy-hn-1",
		Type:			"1",
		Description: 	"no",
		Status:			"running",
		Reserved:		"no",
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