package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/test", getName)

	r.Run(":8080")
}

func getName(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.DefaultQuery("last_name", "last_default_name")

	c.String(http.StatusOK, "%s, %s", firstName, lastName)
}