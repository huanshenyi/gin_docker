package di

import (
	"github.com/gin-gonic/gin"

	"gin_docker/src/domain/authenticator"
	"gin_docker/src/middleware"
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
			gin.Logger(),
			middleware.NewAuth(auth),
			middleware.NewCors(),
			middleware.NewDevice(),
		},
	}
}
