package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("json", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "templategin",
		})
	})

	msg := struct{
		Name string `json:"name"`
		Age int
		Message string
	}{
		"y",
		12,
		"Hello world!",
	}
	r.GET("another", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, msg)
	})

	r.Run(":9091")
}