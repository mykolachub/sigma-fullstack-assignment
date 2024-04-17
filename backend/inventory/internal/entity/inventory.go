package entity

type Inventory struct {
	ID       string
	Name     string
	Quantity int
	Price    int
}

// If some of the value set to true, it will be updated even if no values passed.
// Can be used for resetting inventory
type InventoryForceUpdate struct {
	ID       bool
	Name     bool
	Quantity bool
	Price    bool
}
