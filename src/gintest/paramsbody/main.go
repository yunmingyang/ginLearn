package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()


	r.POST("/test", get)

	r.Run()
}

func get(c *gin.Context) {
	bodyContext, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		c.Abort()
		return
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyContext))
	firstName := c.PostForm("first_name")
	lastName := c.DefaultPostForm("last_name", "default_last_name")

	c.String(http.StatusOK, "%s, %s, %s", lastName, firstName, string(bodyContext))
}