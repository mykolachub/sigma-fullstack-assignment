package service

import (
	"context"
	"errors"
	"fmt"
	"sigma-order/internal/entity"
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
	// TODO: Check if user exists
	// TODO: create config for http requests to user service

	// Create Order
	order, err := s.orderRepo.CreateOrder(entity.Order{UserID: in.UserId, Status: entity.StatusDraft})
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed to create order")
	}

	// Reserve Products
	orderItems := []*pb.OrderItem{}
	for _, v := range in.OrderItems {
		orderItem := pb.OrderItem{ProductId: v.ProductId, Quantity: v.Quantity}
		orderItems = append(orderItems, &orderItem)
	}
	reserved, err := s.invSvc.ReserveInventory(ctx, &pb.ReserveInventoryRequest{OrderItems: orderItems})
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed reserve products")
	}

	// Create Order Items
	products := []*pb.Item{}
	for _, v := range reserved.ReservedProducts {
		product := pb.Item{ProductId: v.ProductId, Quantity: v.Quantity}
		products = append(products, &product)

		_, err := s.orderItemRepo.CreateItem(order.ID, v.ReservedId)
		if err != nil {
			return &pb.OrderResponse{}, errors.New("failed to crete order item")
		}
	}

	res := pb.OrderResponse{Id: order.ID, UserId: order.UserID, OrderItems: products, Status: string(order.Status)}
	return &res, nil
}

func (s *OrderService) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.orderRepo.GetOrder(in.Id)
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed to get order")
	}

	items := []*pb.Item{}
	for _, v := range order.Items {
		reserved, err := s.invSvc.GetReservedInventory(ctx, &pb.GetReservedInventoryRequest{ReservedId: v.ReservedID})
		if err != nil {
			return &pb.OrderResponse{}, errors.New("failed to get reserved item")
		}

		item := pb.Item{ProductId: reserved.ProductId, Quantity: reserved.Quantity, Id: v.ID}
		items = append(items, &item)
	}

	res := pb.OrderResponse{Id: order.ID, UserId: order.UserID, OrderItems: items, Status: string(order.Status)}
	return &res, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, in *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	res := pb.GetAllOrdersResponse{}

	order_ids, err := s.orderRepo.GetOrdersIds()
	if err != nil {
		return &pb.GetAllOrdersResponse{}, errors.New("failed to get orders")
	}

	for _, v := range order_ids {
		order, err := s.GetOrder(ctx, &pb.GetOrderRequest{Id: v})
		if err != nil {
			return &pb.GetAllOrdersResponse{}, errors.New("failed to get orders")
		}
		res.Orders = append(res.Orders, order)
	}

	return &res, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	// TODO: Check if user exists

	order, err := s.orderRepo.UpdateOrder(in.Id, entity.Order{UserID: in.UserId, Status: entity.OrderStatus(in.Status)})
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed to update order")
	}

	res := pb.OrderResponse{Id: order.ID, UserId: order.UserID, Status: string(order.Status)}
	return &res, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, in *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	// TODO: Check if user exists

	order, err := s.orderRepo.GetOrder(in.OrderId)
	if err != nil {
		return &pb.DeleteOrderResponse{}, errors.New("no such order")
	}

	if order.Status != entity.StatusDraft {
		return &pb.DeleteOrderResponse{}, errors.New("order cannot be deleted")
	}

	if len(order.Items) != 0 {
		// Free product reservation
		reserved := []*pb.ReservedProduct{}
		for _, v := range order.Items {
			r := pb.ReservedProduct{ReservedId: v.ReservedID}
			reserved = append(reserved, &r)
		}
		_, err := s.invSvc.FreeReservedInventory(ctx, &pb.FreeReservedInventoryRequest{ReservedProducts: reserved})
		if err != nil {
			return &pb.DeleteOrderResponse{}, err
		}

		// Delete order items
		for _, v := range order.Items {
			_, err := s.orderItemRepo.DeleteItem(v.ID)
			if err != nil {
				return &pb.DeleteOrderResponse{}, errors.New("failed to delete order item")
			}
		}
	}

	// Delete order
	_, err = s.orderRepo.DeleteOrder(order.ID)
	if err != nil {
		return &pb.DeleteOrderResponse{}, errors.New("failed to delete order")
	}

	return &pb.DeleteOrderResponse{OrderId: order.ID}, nil
}

func (s *OrderService) AddOrderItem(ctx context.Context, in *pb.AddOrderItemRequest) (*pb.OrderResponse, error) {
	// TODO: Check if user exists

	// Reserve Product
	orderItem := pb.OrderItem{ProductId: in.OrderItem.ProductId, Quantity: in.OrderItem.Quantity}
	fmt.Printf("in: %+v\n", in)
	reserved, err := s.invSvc.ReserveInventory(ctx, &pb.ReserveInventoryRequest{OrderItems: []*pb.OrderItem{&orderItem}})
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed to reserve product")
	}

	// Create Order Items
	for _, v := range reserved.ReservedProducts {
		_, err := s.orderItemRepo.CreateItem(in.OrderId, v.ReservedId)
		if err != nil {
			return &pb.OrderResponse{}, errors.New("failed to crete order item")
		}
	}

	return s.GetOrder(ctx, &pb.GetOrderRequest{Id: in.OrderId})
}

func (s *OrderService) RemoveOrderItem(ctx context.Context, in *pb.RemoveOrderItemRequest) (*pb.OrderResponse, error) {
	// TODO: Check if user exists

	order, err := s.orderRepo.GetOrder(in.OrderId)
	if err != nil {
		return &pb.OrderResponse{}, err
	}
	for _, v := range order.Items {
		if v.ID == in.OrderItemId {
			// Free Product
			reserved_products := []*pb.ReservedProduct{}
			reserved := pb.ReservedProduct{ReservedId: v.ReservedID}
			reserved_products = append(reserved_products, &reserved)
			_, err := s.invSvc.FreeReservedInventory(ctx, &pb.FreeReservedInventoryRequest{ReservedProducts: reserved_products})
			if err != nil {
				return &pb.OrderResponse{}, err
			}
		}
	}

	// Delete Order Item
	_, err = s.orderItemRepo.DeleteItem(in.OrderItemId)
	if err != nil {
		return &pb.OrderResponse{}, errors.New("failed to delete order item")
	}

	return s.GetOrder(ctx, &pb.GetOrderRequest{Id: in.OrderId})
}

func (s *OrderService) PayOrder(ctx context.Context, in *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	// TODO: Check if user exists

	order, err := s.orderRepo.GetOrder(in.OrderId)
	if err != nil {
		return &pb.PayOrderResponse{}, errors.New("no such order")
	}
	if order.Status == entity.StatusPaid {
		return &pb.PayOrderResponse{}, errors.New("order is already paid")
	}

	// Set status to  IN_PROGRESS
	s.orderRepo.UpdateOrder(in.OrderId, entity.Order{Status: entity.StatusInProgress})

	// Free reserved items
	if len(order.Items) == 0 {
		return &pb.PayOrderResponse{}, errors.New("no items in order")
	}

	// Payment...

	// Set status to PAID
	s.orderRepo.UpdateOrder(in.OrderId, entity.Order{Status: entity.StatusPaid})

	// Decrement items
	reserved_products := []*pb.ReservedProduct{}
	items := []*pb.OrderItem{}
	for _, v := range order.Items {
		reserved_products = append(reserved_products, &pb.ReservedProduct{ReservedId: v.ReservedID})

		reserved, err := s.invSvc.GetReservedInventory(ctx, &pb.GetReservedInventoryRequest{ReservedId: v.ReservedID})
		if err != nil {
			return &pb.PayOrderResponse{}, err
		}

		item := pb.OrderItem{ProductId: reserved.ProductId, Quantity: reserved.Quantity}
		items = append(items, &item)
	}
	_, err = s.invSvc.FreeReservedInventory(ctx, &pb.FreeReservedInventoryRequest{ReservedProducts: reserved_products})
	if err != nil {
		return &pb.PayOrderResponse{}, err
	}

	_, err = s.invSvc.DecrementInventory(ctx, &pb.DecrementInventoryRequest{OrderItems: items})
	if err != nil {
		return &pb.PayOrderResponse{}, err
	}

	return &pb.PayOrderResponse{}, nil
}
