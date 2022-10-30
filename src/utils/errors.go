package utils

import (
	"bytes"
	"fmt"
	"strings"

	"gin_docker/src/log_source"

	"github.com/go-playground/validator/v10"
)

type Errors []error

func (e Errors) Error() string {
	var errors = []string{}
	for _, item := range e {
		errors = append(errors, item.Error())
	}
	return strings.Join(errors, "; ")
}

type SQLError struct {
	GormErrors Errors
}

// ResourceNotFoundError
type ResourceNotFoundError struct {
	Resource string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

// ResourceNotPublicError
type ResourceNotPublicError struct {
	Resource string
}

func (e *ResourceNotPublicError) Error() string {
	return fmt.Sprintf("%s is not public", e.Resource)
}

// NotImplementedYetError
type NotImplementedYetError struct {
	Method string
}

func (e *NotImplementedYetError) Error() string {
	return fmt.Sprintf("%s not implemented yet", e.Method)
}

// UserNotFoundError
type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprint("user not found")
}

// UnauthorizedError
type UnauthorizedError struct {
	Action string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("not authorized to %s", e.Action)
}

// InvalidOutputError
type InvalidOutputError struct {
	Message string
}

func (e *InvalidOutputError) Error() string {
	return e.Message
}

// ConflictError
type ConflictError struct {
	Resource string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("%s conflict", e.Resource)
}

// InvalidTokenError
type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return fmt.Sprintf("invalid token")
}

// InvalidParamError
type InvalidParamError struct {
	Err error
}

func (e *InvalidParamError) Error() string {
	switch e.Err.(type) {
	case validator.ValidationErrors:
		ve := e.Err.(validator.ValidationErrors)
		keys := make([]validator.FieldError, 0, len(ve))
		for _, k := range ve {
			keys = append(keys, k)
		}

		// 以下の処理はvalidatorパッケージと同様
		const (
			blank       = ""
			fieldErrMsg = "Key: '%s' Error:Field validation for '%s' failed on the '%s' tag"
		)
		buff := bytes.NewBufferString(blank)
		for _, key := range keys {
			_, err := buff.WriteString(fmt.Sprintf(fieldErrMsg, key.Namespace(), key.Field(), key.Tag()))
			if err != nil {
				log_source.Log.Error(fmt.Sprintf("cannot write string to buffer: %s", fmt.Sprintf(fieldErrMsg, key.Namespace(), key.Field(), key.Tag())))
				return ""
			}
			_, err = buff.WriteString("\n")
			if err != nil {
				log_source.Log.Error(fmt.Sprintf("cannot write string to buffer"))
				return ""
			}
		}
		return strings.TrimSpace(buff.String())
	default:
		if e.Err != nil {
			return e.Err.Error()
		}

		return ""
	}
}

type DBInternalError struct {
	SQLError
	Err error
}

func (e *DBInternalError) Error() string {
	return appendGormError(fmt.Sprintf("internal error has occurred: %s", e.Err), e.GormErrors)
}

func appendGormError(str string, errs Errors) string {
	return str + fmt.Sprintf("(gorm:%s)", errs.Error())
}
