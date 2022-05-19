package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")

		ctx.String(http.StatusOK, "name is: " + name + ", age is: " + age)
	})

	r.GET("/blog/:year/:month", func(ctx *gin.Context) {
		year := ctx.Param("year")
		month := ctx.Param("month")

		ctx.JSON(http.StatusOK, gin.H{
			"year": year,
			"month": month,
		})
	})

	r.Run()
}