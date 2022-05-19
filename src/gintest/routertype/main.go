package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/get", get)
	r.POST("/post", post)
	r.Handle("DELETE", "/delete", delete)
	r.Any("/any", any)

	r.Run()
}

func get(c *gin.Context) {
	c.String(http.StatusOK, "get")
}

func post(c *gin.Context) {
	c.String(http.StatusOK, "post")
}

func delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}

func any(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}