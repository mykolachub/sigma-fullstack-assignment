package service

import "sigma-inventory/internal/entity"

type Storages struct {
	InventoryRepo InventoryRepo
}

// Interface of methods for working with database
type InventoryRepo interface {
	CreateInventory(data entity.Inventory) (entity.Inventory, error)
	GetInventory(id string, data entity.Inventory) (entity.Inventory, error)
	GetAllInventory(id string, data entity.Inventory) ([]entity.Inventory, error)
	UpdateInventory(id string, data entity.Inventory) (entity.Inventory, error)
	DeleteInventory(id string) (entity.Inventory, error)
}
