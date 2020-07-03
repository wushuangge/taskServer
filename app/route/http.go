package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SetupHttp(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		v1.GET("/task", HandlePosts)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/task", HandleProcess)
	}
}

func HandlePosts(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, posts")
}

func HandleSeries(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, series")
}

func HandleProcess(c *gin.Context)  {
	file, header, err := c.Request.FormFile("uploaded")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err ==nil{
		c.String(http.StatusOK,string(data))
	}
	//文件的名称
	filename := header.Filename
	fmt.Println(file, err, filename)
}