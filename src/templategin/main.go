package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func (s string) template.HTML {
			return template.HTML("<a href=\"" + s + "\">test</a>")
		},
	})
	r.Static("/test", "./statics")
	// r.LoadHTMLFiles("template/index.tmpl")
	r.LoadHTMLGlob("template/**/*")

	// r.GET("index", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Hello",
	// 	})
	// })
	r.GET("posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts", gin.H{
			"title": "1",
		})
	})

	r.GET("users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users", gin.H{
			"title": "2",
		})
	})

	r.Run()
}