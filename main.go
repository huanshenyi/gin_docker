package main

import (
	"gin_docker/src/api/routes"
	"gin_docker/src/di"
	"gin_docker/src/domain"
	"gin_docker/src/infra/db"
	"gin_docker/src/infra/repository"
	"gin_docker/src/log_source"
	"net/http"

	"github.com/sirupsen/logrus"

	_ "gin_docker/src/log_source"

	"github.com/gin-gonic/gin"
)

func main() {

	tx := repository.NewTxEmpty()
	s := di.NewGssktService(tx)
	engine := gin.New()

	// ミドルウェア追加
	mv := di.NewGinMiddlewares(s)
	for _, v := range mv.Chain {
		engine.Use(v)
	}

	// テストルート
	engine.GET("/", func(c *gin.Context) {
		log_source.Log.Info("infoのレベル-")
		log_source.Log.Debug("debugのレベル-")
		// Fieldの設定は各パッケージに使ったほうがいい、init()内ではreturnがないため、意味ない
		log_source.Log.WithField("name", "value").Info("add tag")
		log_source.Log.WithFields(logrus.Fields{
			"event": "value",
			"key":   "key",
		}).Info("Fatal内容...")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "121h",
		})
	})
	initializeApp(tx)
	routes.CreateRoutes(engine, s)

	engine.Run(":3001")
}

func initializeApp(conn domain.Tx) {
	db.MigrateOnlyLocal(conn)
}
