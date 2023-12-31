package middlewares

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unt "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (str string, code int) {

	validate := validator.New()
	uni := unt.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CH")
	err := zh.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		fmt.Println(err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), 500
		}
	}
	return "", 200
}
