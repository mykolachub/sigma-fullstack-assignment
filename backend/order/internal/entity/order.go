package entity

type OrderStatus string

const (
	StatusDraft      OrderStatus = "DRAFT"
	StatusInProgress OrderStatus = "INPROGRESS"
	StatusPaid       OrderStatus = "PAID"
)

type Order struct {
	ID     string
	UserID string
	Status OrderStatus
	Items  []OrderItem
}
