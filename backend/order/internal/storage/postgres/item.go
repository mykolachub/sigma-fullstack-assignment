package postgres

import (
	"database/sql"
	"sigma-order/internal/entity"

	_ "github.com/lib/pq"
)

type OrderItemRepo struct {
	db *sql.DB
}

func InitOrderItemRepo(db *sql.DB) *OrderItemRepo {
	return &OrderItemRepo{db: db}
}

func (r *OrderItemRepo) CreateItem() (entity.OrderItem, error) {
	return entity.OrderItem{}, nil
}

func (r *OrderItemRepo) GetItem(id string) (entity.OrderItem, error) {
	return entity.OrderItem{}, nil
}

func (r *OrderItemRepo) GetAllItems() ([]entity.OrderItem, error) {
	return []entity.OrderItem{}, nil
}

func (r *OrderItemRepo) UpdateItem(id string, data entity.OrderItem) (entity.OrderItem, error) {
	return entity.OrderItem{}, nil
}

func (r *OrderItemRepo) DeleteItem(id string) (entity.OrderItem, error) {
	return entity.OrderItem{}, nil
}
