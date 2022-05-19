package main

import (
	// "fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(ctx *gin.Context) {
		f, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		// dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)
		if err = ctx.SaveUploadedFile(f, dst); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "reveiced!",
		})
	})

	r.Run()
}