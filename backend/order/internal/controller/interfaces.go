package controller

import (
	"context"
	pb "sigma-order/proto"
)

type Services struct {
	OrderService OrderService
}

type OrderService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.OrderResponse, error)
	GetOrder(context.Context, *pb.GetOrderRequest) (*pb.OrderResponse, error)
	GetAllOrders(context.Context, *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error)
	UpdateOrder(context.Context, *pb.UpdateOrderRequest) (*pb.OrderResponse, error)
	DeleteOrder(context.Context, *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	AddOrderItem(context.Context, *pb.AddOrderItemRequest) (*pb.OrderResponse, error)
	RemoveOrderItem(context.Context, *pb.RemoveOrderItemRequest) (*pb.OrderResponse, error)
	PayOrder(context.Context, *pb.PayOrderRequest) (*pb.PayOrderResponse, error)
}
