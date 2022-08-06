package routes

import (
	"gin_docker/src/di"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePublicRoutes(engine *gin.RouterGroup, s *di.GssktService) {
	engine.GET("/tag", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "tag",
		})
	})
}
