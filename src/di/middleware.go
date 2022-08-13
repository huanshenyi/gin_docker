package di

import (
	"gin_docker/src/domain/authenticator"
	"gin_docker/src/middleware"

	"github.com/gin-gonic/gin"
)

type GinMiddlewares struct {
	Chain gin.HandlersChain
}

func NewGinMiddlewares(srv *GssktService) *GinMiddlewares {
	auth := authenticator.NewAuthenticator(
		srv.UserService,
	)
	return &GinMiddlewares{
		Chain: []gin.HandlerFunc{
			gin.Recovery(),
			middleware.NewAuth(auth),
		},
	}
}
