package main

import (
	"gin_test/src/log_source"
	"net/http"

	_ "gin_test/src/log_source"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		log_source.Log.Info("infoのレベル-")
		log_source.Log.Debug("debugのレベル-")
		log_source.Log.WithField("name", "value").Info("add tag")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "h",
		})
	})
	router.Run(":3001")
}
