package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"sigma-inventory/internal/entity"
	"strings"

	_ "github.com/lib/pq"
)

type InventoryRepo struct {
	db *sql.DB
}

func NewInventoryRepo(db *sql.DB) *InventoryRepo {
	return &InventoryRepo{
		db: db,
	}
}

func (r *InventoryRepo) CreateInventory(data entity.Inventory) (entity.Inventory, error) {
	product := entity.Inventory{}

	query := "INSERT INTO Products(name, price, quantity) VALUES($1, $2, $3) RETURNING *"
	rows := r.db.QueryRow(query, data.Name, data.Price, data.Quantity)
	err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return entity.Inventory{}, err
	}
	return product, nil
}

func (r *InventoryRepo) GetInventory(id string) (entity.Inventory, error) {
	product := entity.Inventory{}

	query := "SELECT * FROM products WHERE products.product_id = $1"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return entity.Inventory{}, nil
	}
	return product, nil
}

func (r *InventoryRepo) GetAllInventory() ([]entity.Inventory, error) {
	query := "SELECT * FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return []entity.Inventory{}, err
	}

	products := []entity.Inventory{}
	for rows.Next() {
		product := entity.Inventory{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *InventoryRepo) UpdateInventory(id string, data entity.Inventory, force entity.InventoryForceUpdate) (entity.Inventory, error) {
	product := entity.Inventory{}

	updates := []string{}
	args := []interface{}{id}

	if data.Name != "" || force.Name {
		updates = append(updates, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, data.Name)
	}
	if data.Price != 0 || force.Price {
		updates = append(updates, fmt.Sprintf("price = $%d", len(args)+1))
		args = append(args, data.Price)
	}
	if data.Quantity != 0 || force.Quantity {
		updates = append(updates, fmt.Sprintf("quantity = $%d", len(args)+1))
		args = append(args, data.Quantity)
	}

	if len(updates) == 0 {
		return entity.Inventory{}, errors.New("empty update body")
	}

	query := "UPDATE products SET " + strings.Join(updates, ", ") + " WHERE products.product_id = $1 RETURNING *"
	fmt.Printf("query: %v\n", query)
	err := r.db.QueryRow(query, args...).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return entity.Inventory{}, err
	}

	return product, nil
}

func (r *InventoryRepo) DeleteInventory(id string) (entity.Inventory, error) {
	product := entity.Inventory{}

	query := "DELETE FROM products WHERE products.product_id = $1 RETURNING *"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return entity.Inventory{}, err
	}
	return product, nil
}

func (r *InventoryRepo) ReserveInventory(id string, quantity int) (entity.ReservedInventory, error) {
	reserved := entity.ReservedInventory{}
	product := entity.Inventory{}

	var query string
	var err error

	tx, err := r.db.Begin()
	if err != nil {
		return entity.ReservedInventory{}, err
	}
	defer tx.Rollback()

	query = "SELECT * FROM products WHERE products.product_id = $1"
	err = tx.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		return entity.ReservedInventory{}, err
	}
	if product.Quantity < quantity {
		return entity.ReservedInventory{}, errors.New("not enough product left to reserve")
	}

	// Decrement product quantity
	query = "UPDATE products SET quantity = quantity - $2 WHERE products.product_id = $1 AND quantity - $2 >= 0"
	_, err = tx.Exec(query, id, quantity)
	if err != nil {
		return entity.ReservedInventory{}, err
	}

	query = "INSERT INTO reserved_products(product_id, quantity) VALUES($1, $2) RETURNING *"
	err = tx.QueryRow(query, id, quantity).Scan(&reserved.ID, &reserved.ProductId, &reserved.Quantity)
	if err != nil {
		return entity.ReservedInventory{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.ReservedInventory{}, err
	}

	return reserved, nil
}

func (r *InventoryRepo) FreeReservedInventory(id string) error {
	reserved := entity.ReservedInventory{}

	var query string
	var err error

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query = "SELECT * FROM reserved_products WHERE reserved_id = $1"
	err = tx.QueryRow(query, id).Scan(&reserved.ID, &reserved.ProductId, &reserved.Quantity)
	if err != nil {
		return err
	}

	// Increament product quantity
	query = "UPDATE products SET quantity = quantity + $2 WHERE products.product_id = $1"
	_, err = tx.Exec(query, reserved.ProductId, reserved.Quantity)
	if err != nil {
		return err
	}

	query = "DELETE FROM reserved_products WHERE reserved_id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
