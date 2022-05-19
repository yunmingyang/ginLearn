package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/:name/:id", getInfo)

	r.Run()
}

func getInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"name": c.Param("name"),
			"id": c.Param("id"),
		},
	)
}