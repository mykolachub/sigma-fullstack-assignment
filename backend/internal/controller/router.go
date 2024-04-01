package controller

import (
	"net/http"
	"sigma-test/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, services Services, configs Configs) *gin.Engine {
	// CORS config
	r.Use(cors.New(config.CorsConfig))

	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitUserHandler(r, services.UserService, configs.UserHandlerConfig)
	InitPageHandler(r, services.PageService)

	return r
}
