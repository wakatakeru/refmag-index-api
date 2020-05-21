package infrastructure

import (
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	"github.com/wakatakeru/refmag-index-api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	// Config for CORS (AllowOrigins for development environment)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{
		"Content-Type",
		"Authorization",
	}
	router.Use(cors.New(config))

	paperController := controllers.NewPaperController(NewSqlHandler(), NewJWTHandler())

	router.POST("/papers", func(c *gin.Context) { paperController.Create(c) })
	router.GET("/papers", func(c *gin.Context) { paperController.Index(c) })
	router.GET("/papers/:id", func(c *gin.Context) { paperController.Show(c) })

	Router = router
}
