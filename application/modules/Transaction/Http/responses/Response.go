package responses

import (
	TransactionModels "marketplace-api/application/modules/Transaction/models"
)

type Response struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *TransactionModels.Transaction `json:"data"`
}
