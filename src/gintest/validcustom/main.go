package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)



type Book struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookable" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookable", bookAble)
	}

	r.GET("/bookable", book)

	r.Run()
}

func bookAble(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		tody := time.Now()
		if date.Unix() > tody.Unix() {
			return true
		}
	}
	return false
}

func book(c *gin.Context) {
	var b Book

	if err := c.ShouldBind(&b); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"OK!", "booking":b})
}