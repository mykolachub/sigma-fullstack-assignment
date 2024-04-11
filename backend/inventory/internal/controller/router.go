package controller

import (
	"net/http"
	"sigma-inventory/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Services struct {
	InvService InventoryService
}

func InitRouter(r *gin.Engine, services Services) *gin.Engine {
	// CORS config
	r.Use(cors.New(config.CorsConfig))

	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitInventoryHandler(r, services.InvService)

	return r
}
