package service

import (
	"context"
	"errors"
	"fmt"
	"sigma-inventory/internal/entity"
	"sigma-inventory/proto"
)

type InventoryService struct {
	Repo InventoryRepo
	proto.UnimplementedInventoryServiceServer
}

func NewInventoryService(r InventoryRepo) *InventoryService {
	return &InventoryService{Repo: r}
}

func (s *InventoryService) CreateProduct(ctx context.Context, in *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	data := entity.Inventory{Name: in.Name, Price: int(in.Price), Quantity: int(in.Quantity)}
	product, err := s.Repo.CreateInventory(data)
	if err != nil {
		return &proto.ProductResponse{}, errors.New("failed to create product")
	}

	res := proto.ProductResponse{Id: product.ID, Name: product.Name, Price: int32(product.Price), Quantity: int32(product.Quantity)}
	return &res, nil
}

func (s *InventoryService) GetProduct(ctx context.Context, in *proto.GetProductRequest) (*proto.ProductResponse, error) {
	id := in.Id
	product, err := s.Repo.GetInventory(id)
	if err != nil {
		return &proto.ProductResponse{}, errors.New("no such product")
	}

	res := proto.ProductResponse{Id: product.ID, Name: product.Name, Price: int32(product.Price), Quantity: int32(product.Quantity)}
	return &res, nil
}

func (s *InventoryService) GetProductByReservedId(ctx context.Context, in *proto.GetProductByReservedIdRequest) (*proto.ProductResponse, error) {
	product, err := s.Repo.GetInventoryByReservedId(in.ReservedId)
	if err != nil {
		return &proto.ProductResponse{}, errors.New("failed to get product")
	}

	res := proto.ProductResponse{Id: product.ID, Name: product.Name, Quantity: int32(product.Quantity), Price: int32(product.Price)}
	return &res, nil
}

func (s *InventoryService) GetAllProducts(context.Context, *proto.GetAllProductsRequest) (*proto.GetAllProductsResponse, error) {
	res := []*proto.UpdateProductRequest{}
	products, err := s.Repo.GetAllInventory()
	if err != nil {
		return &proto.GetAllProductsResponse{}, errors.New("failed get products")
	}

	for _, v := range products {
		res = append(res, &proto.UpdateProductRequest{Id: v.ID, Name: v.Name, Quantity: int32(v.Quantity), Price: int32(v.Price)})
	}

	return &proto.GetAllProductsResponse{Products: res}, nil
}

func (s *InventoryService) DecrementInventory(ctx context.Context, in *proto.DecrementInventoryRequest) (*proto.DecrementInventoryResponse, error) {
	for _, v := range in.OrderItems {
		product, err := s.GetProduct(ctx, &proto.GetProductRequest{Id: v.ProductId})
		if err != nil {
			return &proto.DecrementInventoryResponse{}, errors.New("no such product")
		}
		if product.Quantity < v.Quantity {
			return &proto.DecrementInventoryResponse{}, fmt.Errorf("order quantity is greater then product left: %v", v)
		}

		data := entity.Inventory{Quantity: int(product.Quantity) - int(v.Quantity)}
		_, err = s.Repo.UpdateInventory(v.ProductId, data, entity.InventoryForceUpdate{Quantity: true})
		if err != nil {
			return &proto.DecrementInventoryResponse{}, errors.New("failed to decrement inventory")
		}
	}

	return &proto.DecrementInventoryResponse{Success: true}, nil
}

func (s *InventoryService) DeleteProduct(ctx context.Context, in *proto.DeleteProductRequest) (*proto.ProductResponse, error) {
	id := in.Id
	product, err := s.Repo.DeleteInventory(id)
	if err != nil {
		return &proto.ProductResponse{}, errors.New("failed to delete user")
	}

	res := proto.ProductResponse{Id: product.ID, Name: product.Name, Price: int32(product.Price), Quantity: int32(product.Quantity)}
	return &res, nil
}

func (s *InventoryService) ReserveInventory(ctx context.Context, in *proto.ReserveInventoryRequest) (*proto.ReserveInventoryResponse, error) {
	reservedProduct := []*proto.ReservedProduct{}

	for _, v := range in.OrderItems {
		reserved, err := s.Repo.ReserveInventory(v.ProductId, int(v.GetQuantity()))
		if err != nil {
			return &proto.ReserveInventoryResponse{}, errors.New("failed to reserve product")
		}

		reservedProduct = append(reservedProduct, &proto.ReservedProduct{ReservedId: reserved.ID, Quantity: int32(reserved.Quantity), ProductId: reserved.ProductId})
	}
	return &proto.ReserveInventoryResponse{ReservedProducts: reservedProduct}, nil
}

func (s *InventoryService) FreeReservedInventory(ctx context.Context, in *proto.FreeReservedInventoryRequest) (*proto.FreeReservedInventoryResponse, error) {
	for _, v := range in.ReservedProducts {
		err := s.Repo.FreeReservedInventory(v.ReservedId)
		if err != nil {
			return &proto.FreeReservedInventoryResponse{}, errors.New("failed to free product reservation")
		}
	}
	return &proto.FreeReservedInventoryResponse{Success: true}, nil
}

func (s *InventoryService) UpdateProduct(ctx context.Context, in *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	data := entity.Inventory{Name: in.Name, Price: int(in.Price), Quantity: int(in.Quantity)}
	id := in.Id
	product, err := s.Repo.UpdateInventory(id, data, entity.InventoryForceUpdate{})
	if err != nil {
		return &proto.ProductResponse{}, errors.New("failed to update product")
	}

	res := proto.ProductResponse{Id: product.ID, Name: product.Name, Price: int32(product.Price), Quantity: int32(product.Quantity)}
	return &res, nil
}

func (s *InventoryService) GetReservedInventory(ctx context.Context, in *proto.GetReservedInventoryRequest) (*proto.ReservedProduct, error) {
	reserved, err := s.Repo.GetReservedInventory(in.ReservedId)
	if err != nil {
		return &proto.ReservedProduct{}, errors.New("failed to get reserved product")
	}

	res := proto.ReservedProduct{ReservedId: reserved.ID, ProductId: reserved.ProductId, Quantity: int32(reserved.Quantity)}
	return &res, nil
}
