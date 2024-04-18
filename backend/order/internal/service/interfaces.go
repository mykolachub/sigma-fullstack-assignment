package service

import "sigma-order/internal/entity"

type Storages struct {
	OrderRepo     OrderRepo
	OrderItemRepo OrderItemRepo
}

type OrderRepo interface {
	CreateOrder(data entity.Order) (entity.Order, error)
	GetOrder(id string) (entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	UpdateOrder(id string, data entity.Order) (entity.Order, error)
	DeleteOrder(id string) (entity.Order, error)
}

type OrderItemRepo interface {
	CreateItem() (entity.OrderItem, error)
	GetItem(id string) (entity.OrderItem, error)
	GetAllItems() ([]entity.OrderItem, error)
	UpdateItem(id string, data entity.OrderItem) (entity.OrderItem, error)
	DeleteItem(id string) (entity.OrderItem, error)
}
