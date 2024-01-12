package controller

import (
	"net/http"
	"sigma-test/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, services Services) {
	// CORS config
	r.Use(cors.New(config.CorsConfig))

	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitUserHandler(r, services.UserService)

	// Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
