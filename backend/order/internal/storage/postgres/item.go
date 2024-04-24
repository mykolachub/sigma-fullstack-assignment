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

func (r *OrderItemRepo) CreateItem(order_id, reserved_id string) (entity.OrderItem, error) {
	item := entity.OrderItem{}

	query := "INSERT INTO order_items(order_id, reserved_id) VALUES($1, $2) RETURNING *"
	rows := r.db.QueryRow(query, order_id, reserved_id)
	err := rows.Scan(&item.ID, &item.OrderID, &item.ReservedID)
	if err != nil {
		return entity.OrderItem{}, err
	}

	return item, nil
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
	item := entity.OrderItem{}

	query := "DELETE FROM order_items WHERE order_items.item_id = $1 RETURNING *"
	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.OrderID, &item.ReservedID)
	if err != nil {
		return entity.OrderItem{}, err
	}
	return item, nil
}
