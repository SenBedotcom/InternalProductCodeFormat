package domain

type InputOrder struct {
	No                int    `json:"no"`
	PlatformProductID string `json:"platformProductId"`
	Qty               int    `json:"qty"`
	UnitPrice         int    `json:"unitPrice"`
	TotalPrice        int    `json:"totalPrice"`
}

type CleanedOrder struct {
	No         int    `json:"no"`
	ProductID  string `json:"productId"`
	MaterialID string `json:"materialId"`
	ModelID    string `json:"modelId"`
	Qty        int    `json:"qty"`
	UnitPrice  int    `json:"unitPrice"`
	TotalPrice int    `json:"totalPrice"`
}
