package responses

import (
	TransactionModels "marketplace-api/application/modules/Transaction/models"
)

type PaginationResponse struct {
	Status          int                             `json:"status"`
	Message         string                          `json:"message"`
	Data            []TransactionModels.Transaction `json:"data"`
	RecordsFiltered int32                           `json:"records_filtered"`
	RecordsTotal    int32                           `json:"records_total"`
}
