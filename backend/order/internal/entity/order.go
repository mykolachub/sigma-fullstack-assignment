package entity

type orderStatus string

const (
	StatusDraft      orderStatus = "DRAFT"
	StatusInProgress orderStatus = "INPROGRESS"
	StatusPaid       orderStatus = "PAID"
)

type Order struct {
	ID     string
	UserID string
	Status orderStatus
}
