package cmd

import (
	"fmt"
	"log"
	"net"
	"sigma-order/config"
	"sigma-order/internal/controller"
	"sigma-order/internal/service"
	"sigma-order/internal/storage/postgres"
	pb "sigma-order/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(env *config.Env) {

	go grpcServer(env.GrpcPort)

	db, err := postgres.InitDBConnection(postgres.PostgresConfig{
		DBUser:     env.PostgresDBUser,
		DBName:     env.PostgresDBName,
		DBPassword: env.PostgresDBPassword,
		DBPort:     env.PostgresDBPort,
		DBSSLMode:  env.PostgresDBSSLMode,
		DBHost:     env.PostgresDBHost,
	})
	if err != nil {
		log.Fatal(err)
	}

	strorages := service.Storages{
		OrderRepo:     postgres.InitOrderRepo(db),
		OrderItemRepo: postgres.InitOrderItemRepo(db),
	}

	inventoryClient := inventoryClient(env.GrpcInventoryClientHost, env.GrpcInventoryClientPort)
	userServiceBase := fmt.Sprintf("http://%s:%s", env.UserServiceHost, env.UserServicePort)
	userServiceUrl := fmt.Sprintf("%s/api/users/", userServiceBase)
	userService := service.NewUserService(service.UserServiceConfig{Url: userServiceUrl})
	services := controller.Services{
		OrderService: service.NewOrderService(strorages.OrderRepo, strorages.OrderItemRepo, inventoryClient, userService),
	}

	config.InitCorsConfig()
	router := gin.Default()
	r := controller.InitRouter(router, services)

	r.Run(fmt.Sprintf(":%v", env.HttpPort))
}

func grpcServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterOrderServiceServer(s, pb.UnimplementedOrderServiceServer{})
	log.Printf("gRPC server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gRPC failed to server on %v\n", err)
	}
}

func inventoryClient(host, port string) pb.InventoryServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewInventoryServiceClient(conn)
	return client
}
