package responses

type PaginationResponse struct {
	Status          int           `json:"status"`
	Message         string        `json:"message"`
	Data            []interface{} `json:"data"`
	RecordsFiltered int32         `json:"records_filtered"`
	RecordsTotal    int32         `json:"records_total"`
}
