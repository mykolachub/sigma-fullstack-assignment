package service

import (
	"context"
	"sigma-inventory/proto"
)

type Service struct {
	proto.UnimplementedInventoryServiceServer
}

func (s *Service) CreateProduct(ctx context.Context, in *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{Id: "boss", Name: "Amogus", Quantity: 3}, nil
}
