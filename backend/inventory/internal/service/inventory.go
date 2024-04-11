package service

import (
	"context"
	"sigma-inventory/proto"
)

type InventoryService struct {
	repo InventoryRepo
}

func NewInventoryService(r InventoryRepo) *InventoryService {
	return &InventoryService{repo: r}
}

func (s *InventoryService) CreateProduct(ctx context.Context, in *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{}, nil
}

func (s *InventoryService) DecrementInventory(ctx context.Context, in *proto.DecrementInventoryRequest) (*proto.DecrementInventoryResponse, error) {
	return &proto.DecrementInventoryResponse{}, nil
}

func (s *InventoryService) DeleteProduct(ctx context.Context, in *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{}, nil
}

func (s *InventoryService) GetProduct(ctx context.Context, in *proto.GetProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{}, nil
}

func (s *InventoryService) ReserveInventory(ctx context.Context, in *proto.ReserveInventoryRequest) (*proto.ReserveInventoryResponse, error) {
	return &proto.ReserveInventoryResponse{}, nil
}

func (s *InventoryService) UpdateProduct(ctx context.Context, in *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{}, nil
}
