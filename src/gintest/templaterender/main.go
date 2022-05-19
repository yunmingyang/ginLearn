package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	g := gin.Default()
	g.LoadHTMLGlob("template/*")

	g.GET("test", test)

	g.Run()
}

func test(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "test",
		},
	)
}