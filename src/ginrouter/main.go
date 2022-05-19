package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/index", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.Run()
}