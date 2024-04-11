package cmd

import (
	"fmt"
	"log"
	"net"
	"sigma-inventory/config"
	"sigma-inventory/internal/controller"
	"sigma-inventory/internal/service"
	pb "sigma-inventory/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func server(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &service.Service{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func inventoryClient(port string) (*grpc.ClientConn, pb.InventoryServiceClient) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewInventoryServiceClient(conn)
}

func Run(env *config.Env) {
	// gRPC Server
	go server(env.GrpcPort)

	// gRPC Clients
	invConn, invService := inventoryClient(env.GrpcInventoryClientPort)
	defer invConn.Close()

	// HTTP Server
	router := gin.Default()
	services := controller.Services{
		InvService: invService,
	}
	config.InitCorsConfig()
	r := controller.InitRouter(router, services)

	r.Run(fmt.Sprintf(":%v", env.HttpPort))
}
