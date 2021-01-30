package models

import (
	"marketplace-api/application/base"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                   uuid.UUID            `json:"_id"`
	Code                 string               `json:"code"`
	Status               uint8                `json:"status"`
	SubTotal             int                  `json:"sub_total"`
	PriceTotal           int                  `json:"price_total"`
	DiscountPriceAdded   int                  `json:"discount_price_added"`
	DiscountPriceApplied int                  `json:"discount_price_applied"`
	VoucherPriceAdded    int                  `json:"voucher_price_added"`
	VoucherPriceApplied  int                  `json:"voucher_price_applied"`
	TransactionProduct   []TransactionProduct `json:"transaction_product"`
	base.ModelHasAudit
}
