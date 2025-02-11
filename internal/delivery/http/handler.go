package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/orderprocessing/internal/domain"
	"github.com/yourusername/orderprocessing/internal/usecase"
)

type Handler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewHandler(orderUseCase *usecase.OrderUseCase) *Handler {
	return &Handler{
		orderUseCase: orderUseCase,
	}
}

func (h *Handler) ProcessOrder(c *fiber.Ctx) error {
	var input []domain.InputOrder
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	cleanedOrders, err := h.orderUseCase.ProcessOrders(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cleanedOrders)
}
