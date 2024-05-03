package cmd

import (
	"fmt"
	"log"
	"net"
	"sigma-inventory/config"
	"sigma-inventory/internal/controller"
	"sigma-inventory/internal/service"
	"sigma-inventory/internal/storage/postgres"
	pb "sigma-inventory/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func server(port string, repo service.InventoryRepo) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &service.InventoryService{Repo: repo})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Run(env *config.Env) {
	// Database configuration
	db, err := postgres.InitDBConnection(postgres.PostgresConfig{
		DBUser:     env.PostgresDBUser,
		DBName:     env.PostgresDBName,
		DBPassword: env.PostgresDBPassword,
		DBSSLMode:  env.PostgresDBSSLMode,
		DBPort:     env.PostgresDBPort,
		DBHost:     env.PostgresDBHost,
	})
	if err != nil {
		log.Fatal(err)
	}

	storages := service.Storages{
		InventoryRepo: postgres.NewInventoryRepo(db),
	}

	services := controller.Services{
		InventoryService: service.NewInventoryService(storages.InventoryRepo),
	}

	// gRPC Server
	go server(env.GrpcPort, storages.InventoryRepo)

	// HTTP Server
	config.InitCorsConfig()
	router := gin.Default()
	r := controller.InitRouter(router, services)

	r.Run(fmt.Sprintf(":%v", env.HttpPort))
}
