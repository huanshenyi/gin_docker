package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"

	"gin_docker/src/log_source"
	"gin_docker/src/utils"
)

func ErrorResponse(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(toHTTPStatusCode(err), safeErrorMessage(err))
}

// HTTPStatus を Codeに変換
func toHTTPStatusCode(err error) (code int) {
	if err == nil {
		return http.StatusOK
	}
	switch errors.Cause(err).(type) {
	case *utils.NotImplementedYetError:
		code = http.StatusNotImplemented
	case *utils.UserNotFoundError, *utils.UnauthorizedError:
		code = http.StatusUnauthorized
	case *utils.ConflictError:
		code = http.StatusConflict
	case *utils.ResourceNotFoundError, *utils.ResourceNotPublicError:
		code = http.StatusNotFound
	case *utils.InvalidParamError:
		code = http.StatusBadRequest
	case *utils.DBInternalError, *utils.InvalidOutputError:
		code = http.StatusInternalServerError
	default:
		code = http.StatusInternalServerError
	}
	return
}

// ユーザーに見せていいエラーメッセージだけを返す
func safeErrorMessage(err error) (msg map[string]string) {
	if IsPublicError(err) {
		msg = GetErrorResponse(err)
	} else {
		msg = GetErrorResponseReturnNoMessage(err)
	}
	return
}

func IsPublicError(err error) (public bool) {
	switch errors.Cause(err).(type) {
	case
		*utils.UserNotFoundError,
		*utils.UnauthorizedError,
		*utils.ResourceNotFoundError,
		*utils.ResourceNotPublicError:
		public = true
	}
	return
}

// 共通のエラーレスポンスを生成する
func GetErrorResponse(err error) map[string]string {
	return map[string]string{"message": err.Error()}
}

// 共通のエラーレスポンスを生成する、メッセージをLogに保存、クライアントに返さない
func GetErrorResponseReturnNoMessage(err error) map[string]string {
	log_source.Log.Error(err.Error())
	return map[string]string{"message": "error"}
}
