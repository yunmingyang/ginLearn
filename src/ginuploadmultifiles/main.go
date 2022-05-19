package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	// memory limition
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLFiles("index.html")

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
			return
		}

		files := form.File["files"]
		log.Printf("%v\n", files)
		for i, f := range files {
			log.Println("filename is: ", f.Filename)
			dst := fmt.Sprintf("./%s_%d", f.Filename, i)
			log.Println("destination is: ", dst)
			if err := ctx.SaveUploadedFile(f, dst); err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"err": err.Error(),
				})
				return
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})


	r.Run(":9091")
}