package routes

import (
	"gin_docker/src/di"

	"github.com/gin-gonic/gin"
)

func CreatePublicRoutes(engine *gin.RouterGroup, s *di.GssktService) {
	engine.GET("/tags", s.Tag.List)

	user := engine.Group("/user")
	{
		user.POST("/regist", s.User.Regist)
		user.POST("/login", s.User.Login)
	}
	engine.GET("/recruitment", s.Recruitment.PublicList)
}
