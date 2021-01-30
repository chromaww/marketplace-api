package requests

type TransactionCreateRequest struct {
	ProductIds []string `json:"product_ids" binding:"required"`
	Quantities []int16  `json:"quantities" binding:"required"`
}

type TransactionUpdateRequest struct {
	ProductIds []string `json:"product_ids" binding:"required"`
	Quantities []int16  `json:"quantities" binding:"required"`
	Status     uint8    `json:"status" binding:"required"`
}

type TransactionCheckoutRequest struct {
	TransactionId string   `json:"transaction_id" binding:"required"`
	ProductIds    []string `json:"product_ids" binding:"required"`
}

type TransactionDeleteItemRequest struct {
	ItemId string `json:"item_id" binding:"required"`
}
