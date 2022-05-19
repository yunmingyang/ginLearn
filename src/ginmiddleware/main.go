package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)



func indexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"PATH": "/index",
	})
}

func timeCounter(ctx *gin.Context) {
	log.Println("timeCounter in")

	start := time.Now()
	// Next() will run the followed HandlerFunc function recursively
	ctx.Next()
	// time.Sleep(1 * time.Second)
	// If use goroutine in the middleware, and with gin.Context, we should not use it directly
	// We should use ctx.Copy() to avoid some question about concurrent safe
	cost := time.Since(start)
	log.Printf("timeCounter cost: %v\n", cost)

	log.Println("timeCounter out")
}

func m1(ctx *gin.Context) {
	log.Printf("m1 in")

	start := time.Now()
	ctx.Next()
	// Abort() will stop the chain and return
	// ctx.Abort()
	cost := time.Since(start)
	log.Printf("m1 cost: %v\n", cost)

	log.Printf("m1 out")
}

func m2(ctx *gin.Context) {
	log.Printf("m2 in")

	start := time.Now()
	// ctx.Next()
	// time.Sleep(1 * time.Second)
	ctx.Set("PATH", "/")
	cost := time.Since(start)
	log.Printf("m2 cost: %v\n", cost)

	log.Printf("m2 out")
}


func main() {
	// still use the default middleware(Logger, Recovery)
	r := gin.Default()
	// the sequence of the middleware in Use() should be careful
	r.Use(timeCounter, m2)

	// r.GET("/index", timeCounter, m1, m2, indexHandler)
	r.GET("/t", func(ctx *gin.Context) {
		p, ok := ctx.Get("PATH")
		// ok = false

		if !ok {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"msg": "failed",
			})
			// Stop the followed middleware
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"PATH": p,
		})
	}, m1)

	r.Run(":9091")
}