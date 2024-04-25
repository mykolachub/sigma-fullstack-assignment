package controller

import (
	"net/http"
	"sigma-order/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, services Services) *gin.Engine {
	// CORS config
	r.Use(cors.New(config.CorsConfig))

	// Test endpoint
	r.GET("/api/orders/ping/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitOrderHandler(r, services.OrderService)

	return r
}
