package app

import (
	"fmt"
	"log"
	"sigma-user/config"
	"sigma-user/internal/controller"
	"sigma-user/internal/service"
	"sigma-user/internal/storage/aerospike"
	"sigma-user/internal/storage/postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter(env *config.Env) *gin.Engine {
	// Configs initialization
	config.InitCorsConfig()
	postgresConfig := postgres.PostgresConfig{
		DBUser:     env.PostgresDBUser,
		DBName:     env.PostgresDBName,
		DBPassword: env.PostgresDBPassword,
		DBSSLMode:  env.PostgresDBSSLMode,
		DBPort:     env.PostgresDBPort,
		DBHost:     env.PostgresDBHost,
	}
	aerospikeConfig := aerospike.AerospikeConfig{
		Hostname: env.AerospikeHostname,
		Port:     env.AerospikePort,
	}
	userConfig := service.UserConfig{JwtSecret: env.JWTSecret}
	userHandlerConfig := controller.UserHandlerConfig{JwtSecret: env.JWTSecret}

	// Database connection
	postgresDb, err := postgres.InitDBConnection(postgresConfig)
	if err != nil {
		log.Fatal(err)
	}
	aerospikeClient, err := aerospike.InitDBConnection(aerospikeConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Storages initialization
	storages := service.Storages{
		UserRepo: postgres.NewUsersRepo(postgresDb),
		PageRepo: aerospike.NewPageRepo(aerospikeClient, aerospike.AerospikePageConfig{
			Namespace: "test",
			Set:       "pages",
		}),
	}

	// Service initialization
	services := controller.Services{
		UserService: service.NewUserService(storages.UserRepo, userConfig),
		PageService: service.NewPageService(storages.PageRepo),
	}

	// Gin router
	router := gin.Default()

	// Controllers initialization
	r := controller.InitRouter(router, services, controller.Configs{UserHandlerConfig: userHandlerConfig})
	return r
}

func Run(env *config.Env) {
	port := fmt.Sprintf(":%s", env.HttpPort)

	r := SetupRouter(env)
	r.Run(port)
}
