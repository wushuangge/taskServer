package route

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func SetupHttp(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		v1.GET("/task", HandleGets)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/task", HandlePosts)
	}
}

func HandleGets(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, gets")
}

func HandlePosts(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, posts")
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