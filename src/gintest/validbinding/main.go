package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type Person struct {
	// In binding, "|" means "or", "," means "and"
	Name string `form:"name" binding:"required"`
	Age int `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("testing", testing)

	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	defer c.Request.Body.Close()

	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusBadRequest, "struct bind failed: %v", err)
		c.Abort()
		return
	}

	c.String(http.StatusOK, "get: %v", person)
}