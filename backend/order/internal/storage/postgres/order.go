package postgres

import (
	"database/sql"
	"sigma-order/internal/entity"

	_ "github.com/lib/pq"
)

type OrderRepo struct {
	db *sql.DB
}

func InitOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(data entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (r *OrderRepo) GetOrder(id string) (entity.Order, error) {
	return entity.Order{}, nil
}

func (r *OrderRepo) GetAllOrders() ([]entity.Order, error) {
	return []entity.Order{}, nil
}

func (r *OrderRepo) UpdateOrder(id string, data entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (r *OrderRepo) DeleteOrder(id string) (entity.Order, error) {
	return entity.Order{}, nil
}
