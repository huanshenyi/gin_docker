package controller

import (
	"time"

	"gin_docker/src/utils"

	"github.com/go-playground/validator/v10"
)

var valid validator.Validate

func init() {
	valid = *validator.New()

	if err := valid.RegisterValidation("dateFormat", dateFormat); err != nil {
		panic(err)
	}
}

func Validate(input interface{}) (err error) {
	if err = valid.Struct(input); err != nil {
		return &utils.InvalidParamError{Err: err}
	}
	return
}

func dateFormat(fl validator.FieldLevel) bool {
	if _, err := time.Parse(fl.Param(), fl.Field().String()); err != nil {
		return false
	}
	return true
}
