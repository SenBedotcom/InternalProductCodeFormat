package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/orderprocessing/internal/delivery/http"
	"github.com/yourusername/orderprocessing/internal/repository/memory"
	"github.com/yourusername/orderprocessing/internal/usecase"
)

func main() {
	app := fiber.New()

	// Initialize repositories
	orderRepo := memory.NewOrderRepository()

	// Initialize use cases
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Initialize handlers
	handler := http.NewHandler(orderUseCase)

	// Setup routes
	api := app.Group("/api")
	api.Post("/orders/process", handler.ProcessOrder)

	log.Fatal(app.Listen(":3000"))
}
