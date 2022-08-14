package middleware

import (
	"fmt"
	"gin_docker/src/domain"
	"gin_docker/src/infra"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// クロスドメイン解決用
func NewCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	fmt.Println(infra.EnvMan.CORSAllowOrigins)
	config.AllowOrigins = infra.EnvMan.CORSAllowOrigins
	config.AllowMethods = []string{"GET", "PUT", "POST", "PATCH", "HEAD", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{domain.ClientAccessUserTokenHTTPHeaderKey}
	return cors.New(
		config,
	)
}
