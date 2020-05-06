package infrastructure

import (
	gin "github.com/gin-gonic/gin"
	"github.com/wakatakeru/refmag-index-api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	paperController := controllers.NewPaperController(NewSqlHandler())

	router.POST("/papers", func(c *gin.Context) { paperController.Create(c) })
	router.GET("/papers", func(c *gin.Context) { paperController.Index(c) })

	Router = router
}
