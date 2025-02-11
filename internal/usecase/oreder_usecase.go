package usecase

import (
	"github.com/yourusername/orderprocessing/internal/domain"
	"github.com/yourusername/orderprocessing/pkg/orderparser"
)

type OrderUseCase struct {
	repo domain.OrderRepository
}

func NewOrderUseCase(repo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
	}
}

func (u *OrderUseCase) ProcessOrders(input []domain.InputOrder) ([]domain.CleanedOrder, error) {
	parser := orderparser.NewParser()
	cleanedOrders := parser.ParseOrders(input)

	if err := u.repo.SaveOrder(cleanedOrders); err != nil {
		return nil, err
	}

	return cleanedOrders, nil
}
