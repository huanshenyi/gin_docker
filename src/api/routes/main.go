package routes

import (
    "gin_test/src/di"
    "github.com/gin-gonic/gin"
)

func CreateRoutes(engine *gin.Engine, s *di.GssktService) {
    public := engine.Group("/public")
    private := engine.Group("/private")
    CreatePublicRoutes(public, s)
    CreatePrivateRoutes(private, s)
}