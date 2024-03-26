package app

import (
	"database/sql"
	"fmt"
	"log"
	"sigma-test/config"
	"sigma-test/internal/controller"
	"sigma-test/internal/service"
	"sigma-test/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	env := config.ConfigEnv()
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", env.DBUser, env.DBName, env.DBPassword, env.DBSSLMode)
	db, err := sql.Open("postgres", connStr)
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
