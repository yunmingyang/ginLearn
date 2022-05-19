package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)



func main() {
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatalf("file create err: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, f)

	g := gin.New()
	// gin.Recovery will prevent the web server dying from the panic
	g.Use(gin.Logger(), gin.Recovery())

	g.GET("/middleware", middleware)

	g.Run()
}

func middleware(c *gin.Context) {
	name := c.DefaultQuery("name", "defaut_name")

	// panic("panic to test gin.Recovery")
	c.String(http.StatusOK, name)
}