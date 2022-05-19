package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"01/02/2006"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	defer c.Request.Body.Close()

	// accroding to the content type of the request to handle different bind operations
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusBadRequest, "struct bind failed: %v", err.Error())
		c.Abort()
		return
	}
	c.String(http.StatusOK, "%v", person)
}