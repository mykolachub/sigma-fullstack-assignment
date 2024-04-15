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
	DeleteProduct(context.Context, *proto.DeleteProductRequest) (*proto.ProductResponse, error)
	FreeReservedInventory(context.Context, *proto.FreeReservedInventoryRequest) (*proto.FreeReservedInventoryResponse, error)
	GetProduct(context.Context, *proto.GetProductRequest) (*proto.ProductResponse, error)
	GetAllProducts(context.Context, *proto.GetAllProductsRequest) (*proto.GetAllProductsResponse, error)
	ReserveInventory(context.Context, *proto.ReserveInventoryRequest) (*proto.ReserveInventoryResponse, error)
	UpdateProduct(context.Context, *proto.UpdateProductRequest) (*proto.ProductResponse, error)
}
