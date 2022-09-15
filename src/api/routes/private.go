package routes

import (
	"github.com/gin-gonic/gin"

	"gin_docker/src/di"
)

func CreatePrivateRoutes(engine *gin.RouterGroup, s *di.GssktService) {
	user := engine.Group("/user")
	{
		me := user.Group("/me")
		{
			recruitment := me.Group("/recruitments")
			{
				recruitment.GET("", s.Recruitment.List)
				recruitment.POST("", s.Recruitment.Create)
			}

			joinRecruitment := me.Group("/join-recruitments")
			{
				joinRecruitment.GET("", s.Recruitment.JoinList)
				joinRecruitment.POST("", s.Recruitment.Join)
			}

			profile := me.Group("/profile")
			{
				profile.GET("", s.User.GetMyInfo)
				profile.PUT("", s.User.UpdateMyInfo)
			}
		}
	}
}
