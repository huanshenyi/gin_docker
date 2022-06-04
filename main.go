package main

import (
	"gin_test/src/api/routes"
	"gin_test/src/di"
	"gin_test/src/log_source"
	"github.com/sirupsen/logrus"
	"net/http"

	_ "gin_test/src/log_source"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		log_source.Log.Info("infoのレベル-")
		log_source.Log.Debug("debugのレベル-")
		// Fieldの設定は各パッケージに使ったほうがいい、init()内ではreturnがないため、意味ない
		log_source.Log.WithField("name", "value").Info("add tag")
		log_source.Log.WithFields(logrus.Fields{
			"event":"value",
			"key":"key",
		}).Info("Fatal内容...")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "h1",
		})
	})
	s := di.NewGssktService()
	routes.CreateRoutes(engine, s)

	engine.Run(":3001")
}
