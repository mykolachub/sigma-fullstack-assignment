package controller

import (
	"net/http"
	"sigma-test/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
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
	InitPageHandler(r, services.PageService, gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "track page circuit breaker",
		Timeout: time.Millisecond,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 2
		},
	}))

	return r
}
