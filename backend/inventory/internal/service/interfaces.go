package service

import "sigma-inventory/internal/entity"

type Storages struct {
	InventoryRepo InventoryRepo
}

// Interface of methods for working with database
type InventoryRepo interface {
	CreateInventory(data entity.Inventory) (entity.Inventory, error)
	GetInventory(id string) (entity.Inventory, error)
	GetInventoryByReservedId(reserved_id string) (entity.Inventory, error)
	GetAllInventory() ([]entity.Inventory, error)
	UpdateInventory(id string, data entity.Inventory, force entity.InventoryForceUpdate) (entity.Inventory, error)
	DeleteInventory(id string) (entity.Inventory, error)
	ReserveInventory(id string, quantity int) (entity.ReservedInventory, error)
	FreeReservedInventory(id string) error
	GetReservedInventory(id string) (entity.ReservedInventory, error)
}
