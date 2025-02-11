package domain

type OrderRepository interface {
	SaveOrder(order []CleanedOrder) error
	GetOrders() ([]CleanedOrder, error)
}
