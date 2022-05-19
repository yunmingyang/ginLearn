package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales"
	enlocales "github.com/go-playground/locales/en"
	zhlocales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entrans "github.com/go-playground/validator/v10/translations/en"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
)



var (
	// Validator
	validate *validator.Validate = validator.New()
	// Translator
	zh locales.Translator = zhlocales.New()
	en locales.Translator = enlocales.New()
	uni *ut.UniversalTranslator = ut.New(zh, en)
)

type Person struct {
	Name string `form:"name" validate:"required"`
	Age int `form:"age" validate:"required,gt=10"`
	Address string `form:"address" validate:"required"`
}

func main() {
	r := gin.Default()

	r.GET("testing", testing)

	r.Run(":8080")
}

func testing(c *gin.Context) {
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := uni.GetTranslator(locale)
	switch locale {
	case "zh":
		zhtrans.RegisterDefaultTranslations(validate, trans)
	case "en":
		entrans.RegisterDefaultTranslations(validate, trans)
	default:
		zhtrans.RegisterDefaultTranslations(validate, trans)
	}

	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(
			http.StatusBadGateway,
			"%v",
			err,
		)
		c.Abort()
		return
	}
	if err := validate.Struct(person); err != nil {
		errs := err.(validator.ValidationErrors)
		errSlice := []string{}
		for _, e := range errs {
			errSlice = append(errSlice, e.Translate(trans))
		}
		c.String(http.StatusBadRequest, "%v", errSlice)
		c.Abort()
		return
	}

	c.String(
		http.StatusOK,
		"%v",
		person,
	)
}