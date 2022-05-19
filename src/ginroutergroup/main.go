package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	g := r.Group("/group")
	{
		g.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"PATH": "/group/index",
			})
		})
		g.POST("/insert", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"PATH": "/group/insert",
			})
		})

		m := g.Group("/m")
		{
			m.GET("index", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"PATH": "/group/m/index",
				})
			})
		}
	}

	r.Any("any", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case "GET":
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET",
			})
		case "POST":
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST",
			})
		}
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"PATH": "No route",
		})
	})

	r.Run(":9091")
}