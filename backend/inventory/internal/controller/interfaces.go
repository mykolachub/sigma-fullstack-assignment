package controller

import (
	"context"
	"sigma-inventory/proto"
)

type Services struct {
	InventoryService InventoryService
}

type InventoryService interface {
	CreateProduct(context.Context, *proto.CreateProductRequest) (*proto.ProductResponse, error)
	DecrementInventory(context.Context, *proto.DecrementInventoryRequest) (*proto.DecrementInventoryResponse, error)
	DeleteProduct(context.Context, *proto.UpdateProductRequest) (*proto.ProductResponse, error)
	GetProduct(context.Context, *proto.GetProductRequest) (*proto.ProductResponse, error)
	ReserveInventory(context.Context, *proto.ReserveInventoryRequest) (*proto.ReserveInventoryResponse, error)
	UpdateProduct(context.Context, *proto.UpdateProductRequest) (*proto.ProductResponse, error)
}
