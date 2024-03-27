package app

import (
	"log"
	"sigma-test/config"
	"sigma-test/internal/controller"
	"sigma-test/internal/service"
	"sigma-test/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db, err := postgres.InitDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Storages initialization
	storages := service.Storages{
		UserRepo: postgres.NewUsersRepo(db),
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
	r := controller.InitRouter(router, services)
	return r
}

func Run() {
	env := config.ConfigEnv()
	port := ":" + env.Port

	r := SetupRouter()
	r.Run(port)
}
