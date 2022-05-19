package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	g := gin.Default()
	g.Use(IPAuthMiddleware())

	g.GET("test", test)

	g.Run()
}

func test(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

func IPAuthMiddleware() gin.HandlerFunc {
	return test1
}

func test1(c *gin.Context) {
	ipList := []string{
		"127.0.0.1",
	}
	flag := false

	clientIP := c.ClientIP()
	for _, IP := range ipList {
		if IP == clientIP {
			flag = true
			break
		}
	}

	if !flag {
		c.String(http.StatusUnauthorized, "%s is not in white list", clientIP)
		c.Abort()
	}
}