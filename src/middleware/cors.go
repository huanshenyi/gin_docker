package middleware

import (
	"gin_docker/src/domain"
	"gin_docker/src/infra"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// クロスドメイン解決用
func NewCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = infra.EnvMan.CORSAllowOrigins
	config.AllowMethods = []string{"GET", "PUT", "POST", "PATCH", "HEAD", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{domain.ClientAccessUserTokenHTTPHeaderKey, domain.ClientAppTypeHTTPHeaderKey, "Content-Type"}
	return cors.New(
		config,
	)
}
