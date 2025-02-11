package memory

import (
	"sync"

	"github.com/yourusername/orderprocessing/internal/domain"
)

type orderRepository struct {
	orders []domain.CleanedOrder
	mutex  sync.RWMutex
}

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{
		orders: make([]domain.CleanedOrder, 0),
	}
}

func (r *orderRepository) SaveOrder(orders []domain.CleanedOrder) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.orders = append(r.orders, orders...)
	return nil
}

func (r *orderRepository) GetOrders() ([]domain.CleanedOrder, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.orders, nil
}
