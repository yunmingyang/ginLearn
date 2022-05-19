package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/query", func(ctx *gin.Context) {
		name := ctx.Query("name")

		ctx.String(http.StatusOK, name)
	})

	r.Run()
}