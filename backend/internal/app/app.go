package app

import (
	"fmt"
	"log"
	"sigma-test/config"
	"sigma-test/internal/controller"
	"sigma-test/internal/service"
	"sigma-test/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter(env *config.Env) *gin.Engine {
	// Configs initialization
	config.InitCorsConfig()
	dbConfig := postgres.DbConfig{
		DBUser:     env.DBUser,
		DBName:     env.DBName,
		DBPassword: env.DBPassword,
		DBSSLMode:  env.DBSSLMode,
	}
	userConfig := service.UserConfig{JwtSecret: env.JWTSecret}
	userHandlerConfig := controller.UserHandlerConfig{JwtSecret: env.JWTSecret}

	// Database connection
	db, err := postgres.InitDBConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Storages initialization
	storages := service.Storages{
		UserRepo: postgres.NewUsersRepo(db),
	}

	// Service initialization
	services := controller.Services{
		UserService: service.NewUserService(storages.UserRepo, userConfig),
	}

	// Gin router
	router := gin.Default()

	// Controllers initialization
	r := controller.InitRouter(router, services, controller.Configs{UserHandlerConfig: userHandlerConfig})
	return r
}

func Run(env *config.Env) {
	port := fmt.Sprintf(":%s", env.Port)

	r := SetupRouter(env)
	r.Run(port)
}
