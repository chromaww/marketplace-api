package models

import (
	"marketplace-api/application/base"
	ProductModels "marketplace-api/application/modules/Product/models"

	"github.com/google/uuid"
)

type TransactionProduct struct {
	ID            uuid.UUID             `json:"_id"`
	TransactionID uuid.UUID             `json:"transaction_id"`
	SubTotal      int                   `json:"sub_total"`
	PriceTotal    int                   `json:"price_total"`
	ProductID     uuid.UUID             `json:"product_id"`
	Quantity      uint8                 `json:"quantity"`
	Status        uint8                 `json:"status"`
	Product       ProductModels.Product `gorm:"embedded;embeddedPrefix:product_"`
	base.ModelHasAudit
}
