package controllers

import (
	"net/http"
	"time"

	"marketplace-api/application/base"
	"marketplace-api/application/modules/Transaction/Http/responses"
	"marketplace-api/application/utils/constants"

	TransactionModels "marketplace-api/application/modules/Transaction/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//TransactionRead : Read transaction data
func (this *TransactionController) Read(c *gin.Context) {
	transactionID := c.Param("transactionID")
	var data []TransactionModels.Transaction

	if transactionID != "" {
		trxID, err := uuid.Parse(transactionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid ID",
				Data:    nil,
			})

			return
		}

		data = []TransactionModels.Transaction{
			TransactionModels.Transaction{
				ID:                   trxID,
				Code:                 "TRX0001",
				Status:               constants.ORDER_STATUS.New,
				SubTotal:             10000,
				PriceTotal:           5000,
				DiscountPriceAdded:   5000,
				DiscountPriceApplied: 5000,
				VoucherPriceAdded:    0,
				VoucherPriceApplied:  0,
				TransactionProduct:   make([]TransactionModels.TransactionProduct, 0),
				ModelHasAudit: base.ModelHasAudit{
					CreatedAt: time.Now(),
					UpdatedAt: nil,
					DeletedAt: nil,
				},
			},
		}

		c.JSON(http.StatusOK, responses.PaginationResponse{
			Status:          http.StatusOK,
			Message:         "OK",
			Data:            data,
			RecordsFiltered: 1,
			RecordsTotal:    1,
		})
	} else {
		data = []TransactionModels.Transaction{
			TransactionModels.Transaction{
				ID:                   uuid.New(),
				Code:                 "TRX0001",
				Status:               constants.ORDER_STATUS.New,
				SubTotal:             10000,
				PriceTotal:           5000,
				DiscountPriceAdded:   5000,
				DiscountPriceApplied: 5000,
				VoucherPriceAdded:    0,
				VoucherPriceApplied:  0,
				TransactionProduct:   make([]TransactionModels.TransactionProduct, 0),
				ModelHasAudit: base.ModelHasAudit{
					CreatedAt: time.Now(),
					UpdatedAt: nil,
					DeletedAt: nil,
				},
			},
			TransactionModels.Transaction{
				ID:                   uuid.New(),
				Code:                 "TRX0002",
				Status:               constants.ORDER_STATUS.New,
				SubTotal:             20000,
				PriceTotal:           13000,
				DiscountPriceAdded:   7000,
				DiscountPriceApplied: 7000,
				VoucherPriceAdded:    0,
				VoucherPriceApplied:  0,
				TransactionProduct:   make([]TransactionModels.TransactionProduct, 0),
				ModelHasAudit: base.ModelHasAudit{
					CreatedAt: time.Now(),
					UpdatedAt: nil,
					DeletedAt: nil,
				},
			},
		}

		c.JSON(http.StatusOK, responses.PaginationResponse{
			Status:          http.StatusOK,
			Message:         "OK",
			Data:            data,
			RecordsFiltered: 2,
			RecordsTotal:    2,
		})
	}
}
