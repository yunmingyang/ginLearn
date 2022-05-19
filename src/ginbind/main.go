package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(ctx *gin.Context) {
		var u UserInfo

		if err := ctx.ShouldBind(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "ShouldBind failed: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"usename": u.Username,
			"password": u.Password,
		})
	})

	r.Run(":9091")
}