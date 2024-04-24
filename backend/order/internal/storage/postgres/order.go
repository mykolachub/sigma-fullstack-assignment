package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"sigma-order/internal/entity"
	"strings"

	_ "github.com/lib/pq"
)

type OrderRepo struct {
	db *sql.DB
}

func InitOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(data entity.Order) (entity.Order, error) {
	order := entity.Order{}

	query := "INSERT INTO Orders(user_id, status) VALUES($1, $2) RETURNING *"
	rows := r.db.QueryRow(query, data.UserID, data.Status)
	err := rows.Scan(&order.ID, &order.UserID, &order.Status)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r *OrderRepo) GetOrder(id string) (entity.Order, error) {
	order := entity.Order{}

	var query string
	var err error

	tx, err := r.db.Begin()
	if err != nil {
		return entity.Order{}, err
	}
	defer tx.Rollback()

	// Get info about order
	query = "SELECT * FROM orders WHERE order_id = $1"
	err = tx.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Status)
	if err != nil {
		return entity.Order{}, err
	}

	// Get all order items
	query = "SELECT * FROM order_items WHERE order_items.order_id = $1"
	rows, err := tx.Query(query, id)
	if err != nil {
		return entity.Order{}, err
	}

	items := []entity.OrderItem{}
	for rows.Next() {
		item := entity.OrderItem{}
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ReservedID); err != nil {
			return entity.Order{}, err
		}
		items = append(items, item)
	}
	order.Items = items

	if err := tx.Commit(); err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r *OrderRepo) GetAllOrders() ([]entity.Order, error) {
	orders := []entity.Order{}

	order_ids, err := r.GetOrdersIds()
	if err != nil {
		return []entity.Order{}, nil
	}

	for _, v := range order_ids {
		order, err := r.GetOrder(v)
		if err != nil {
			return []entity.Order{}, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepo) GetOrdersIds() ([]string, error) {
	order_ids := []string{}

	query := "SELECT order_id FROM orders"
	rows, err := r.db.Query(query)
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var order_id string
		if err := rows.Scan(&order_id); err != nil {
			return []string{}, err
		}
		order_ids = append(order_ids, order_id)
	}

	return order_ids, nil
}

func (r *OrderRepo) UpdateOrder(id string, data entity.Order) (entity.Order, error) {
	order := entity.Order{}

	updates := []string{}
	args := []interface{}{id}

	if data.UserID != "" {
		updates = append(updates, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, data.UserID)
	}
	if data.Status != "" {
		updates = append(updates, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, data.Status)
	}

	if len(updates) == 0 {
		return entity.Order{}, errors.New("empty uppdate body")
	}

	query := "UPDATE orders SET " + strings.Join(updates, ", ") + " WHERE orders.order_id = $1 RETURNING *"
	err := r.db.QueryRow(query, args...).Scan(&order.ID, &order.UserID, &order.Status)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r *OrderRepo) DeleteOrder(id string) (entity.Order, error) {
	order := entity.Order{}

	query := "DELETE FROM orders WHERE order_id = $1 RETURNING *"
	err := r.db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Status)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
