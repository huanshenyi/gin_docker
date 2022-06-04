package routes

import (
	"gin_test/src/di"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePrivateRoutes(engine *gin.RouterGroup, s *di.GssktService) {
	engine.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})
}
