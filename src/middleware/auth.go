package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin_docker/src/domain"
	"gin_docker/src/domain/authenticator"
	"gin_docker/src/log_source"
)

const (
	authHeader = domain.ClientAccessUserTokenHTTPHeaderKey
)

func NewAuth(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.GetHeader(authHeader)
		// 未ログインの場合、AnonymousUserをセット
		if v == "" {
			authenticator.SetUser(c, auth.GetAnonymousUser())
			return
		}

		token, tokenType, err := authenticator.GetTokenFromHeader(v)
		if err != nil {
			log_source.Log.Error(fmt.Sprintf("invalid auth header value: %s: %s", authHeader, v))
			authenticator.SetUser(c, auth.GetAnonymousUser())
			return
		}

		user, err := auth.GetUserFromToken(token, tokenType)
		if err != nil {
			log_source.Log.Error(fmt.Sprintf("failed to get user from token: %s", err))
			authenticator.SetUser(c, auth.GetAnonymousUser())
			return
		}

		authenticator.SetAuthedUser(c, user)
	}
}
