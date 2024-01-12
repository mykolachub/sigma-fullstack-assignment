package app

import (
	"sigma-test/config"
	"sigma-test/internal/controller"
	"sigma-test/internal/service"
	"sigma-test/internal/storage/inmemory"

	"github.com/gin-gonic/gin"
)

func Run( /* additional configs */ ) {

	// Storages initialization
	storages := service.Storages{
		UserRepo: inmemory.NewUsersRepo( /* database connection */ ),
	}

	// Service initialization
	services := controller.Services{
		UserService: service.NewUserService(storages.UserRepo),
	}

	// CORS config
	config.InitCorsConfig()

	// Gin router
	router := gin.Default()

	// Controllers initialization
	controller.InitRouter(router, services)
}
