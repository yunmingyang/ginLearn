package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)



func main() {
	g := gin.Default()

	g.GET("test", test)

	h := &http.Server{
		Addr: ":8080",
		Handler: g,
	}
	go func ()  {
		if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen err: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := h.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown err: %s\n", err)
	}

	log.Println("server exist")
}

func test(c *gin.Context) {
	time.Sleep(10 * time.Second)
	c.String(http.StatusOK, "hello")
}