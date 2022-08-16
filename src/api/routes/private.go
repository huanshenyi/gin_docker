package routes

import (
	"gin_docker/src/di"

	"github.com/gin-gonic/gin"
)

func CreatePrivateRoutes(engine *gin.RouterGroup, s *di.GssktService) {
	user := engine.Group("/user")
	{
		me := user.Group("/me")
		{
			recruitment := me.Group("/recruitments")
			{
				recruitment.GET("", s.Recruitment.List)
			}
		}
	}
}
