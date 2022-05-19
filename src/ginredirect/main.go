package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "index")
	})

	r.GET("/redirect_external", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/redirect_internal", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/index"
		r.HandleContext(ctx)
	})

	r.Run(":9090")
}