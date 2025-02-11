package orderparser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/yourusername/orderprocessing/internal/domain"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseOrders(input []domain.InputOrder) []domain.CleanedOrder {
	var result []domain.CleanedOrder
	nextNo := 1

	for _, order := range input {
		// Clean and split the platform product ID
		products := p.splitProducts(order.PlatformProductID)

		for _, product := range products {
			// Extract quantity multiplier if exists
			qty, productID := p.extractQuantity(product)
			if qty == 0 {
				qty = 1
			}

			// Clean product ID
			cleanProductID := p.cleanProductID(productID)

			// Extract material ID and model ID
			materialID, modelID := p.extractIDs(cleanProductID)

			// Calculate unit price
			unitPrice := order.TotalPrice / (qty * len(products))

			// Add main product
			result = append(result, domain.CleanedOrder{
				No:         nextNo,
				ProductID:  cleanProductID,
				MaterialID: materialID,
				ModelID:    modelID,
				Qty:        qty,
				UnitPrice:  unitPrice,
				TotalPrice: unitPrice * qty,
			})
			nextNo++

			// Add complementary items
			result = append(result, p.createComplementaryItems(materialID, qty, &nextNo)...)
		}
	}

	return result
}

func (p *Parser) splitProducts(productID string) []string {
	// Remove any leading/trailing special characters and split by "/"
	cleaned := regexp.MustCompile(`^[^a-zA-Z0-9]+|[^a-zA-Z0-9]+$`).ReplaceAllString(productID, "")
	return strings.Split(cleaned, "/")
}

func (p *Parser) extractQuantity(product string) (int, string) {
	re := regexp.MustCompile(`\*(\d+)`)
	matches := re.FindStringSubmatch(product)
	if len(matches) > 1 {
		qty, _ := strconv.Atoi(matches[1])
		return qty, re.ReplaceAllString(product, "")
	}
	return 0, product
}

func (p *Parser) cleanProductID(productID string) string {
	// Remove any special characters and trim spaces
	cleaned := regexp.MustCompile(`[^a-zA-Z0-9-]+`).ReplaceAllString(productID, "")
	return strings.TrimSpace(cleaned)
}

func (p *Parser) extractIDs(productID string) (string, string) {
	parts := strings.Split(productID, "-")
	if len(parts) < 3 {
		return "", ""
	}

	materialID := parts[0] + "-" + parts[1]
	modelID := strings.Join(parts[2:], "-")

	return materialID, modelID
}

func (p *Parser) createComplementaryItems(materialID string, qty int, nextNo *int) []domain.CleanedOrder {
	var items []domain.CleanedOrder

	// Add wiping cloth
	items = append(items, domain.CleanedOrder{
		No:         *nextNo,
		ProductID:  "WIPING-CLOTH",
		Qty:        qty,
		UnitPrice:  0,
		TotalPrice: 0,
	})
	*nextNo++

	// Add cleaner based on texture
	texture := strings.Split(materialID, "-")[1]
	items = append(items, domain.CleanedOrder{
		No:         *nextNo,
		ProductID:  texture + "-CLEANNER",
		Qty:        qty,
		UnitPrice:  0,
		TotalPrice: 0,
	})
	*nextNo++

	return items
}
