package service

import (
	"context"
	pb "sigma-order/proto"
)

type OrderService struct {
	orderRepo     OrderRepo
	orderItemRepo OrderItemRepo

	invSvc pb.InventoryServiceClient
}

func NewOrderService(orderRepo OrderRepo, orderItemRepo OrderItemRepo, invSvc pb.InventoryServiceClient) *OrderService {
	return &OrderService{orderRepo: orderRepo, orderItemRepo: orderItemRepo, invSvc: invSvc}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) GetOrder(context.Context, *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) GetAllOrders(context.Context, *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	return &pb.GetAllOrdersResponse{}, nil
}

func (s *OrderService) UpdateOrder(context.Context, *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) DeleteOrder(context.Context, *pb.DeleteOrderRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) AddOrderItem(context.Context, *pb.AddOrderItemRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) RemoveOrderItem(context.Context, *pb.RemoveOrderItemRequest) (*pb.OrderResponse, error) {
	return &pb.OrderResponse{}, nil
}

func (s *OrderService) PayOrder(context.Context, *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	return &pb.PayOrderResponse{}, nil
}
