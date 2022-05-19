package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/user/*action", getUserAction)

	r.Run()
}

func getUserAction(c *gin.Context) {
	c.String(http.StatusOK, "hello world!")
}