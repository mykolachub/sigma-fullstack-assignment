package postgres

import (
	"database/sql"
	"sigma-inventory/internal/entity"

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
	return entity.Inventory{}, nil
}

func (r *InventoryRepo) GetInventory(id string, data entity.Inventory) (entity.Inventory, error) {
	return entity.Inventory{}, nil
}

func (r *InventoryRepo) GetAllInventory(id string, data entity.Inventory) ([]entity.Inventory, error) {
	return []entity.Inventory{}, nil
}

func (r *InventoryRepo) UpdateInventory(id string, data entity.Inventory) (entity.Inventory, error) {
	return entity.Inventory{}, nil
}

func (r *InventoryRepo) DeleteInventory(id string) (entity.Inventory, error) {
	return entity.Inventory{}, nil
}
