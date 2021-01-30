package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Read: Read transaction data
func (this *TransactionController) Read(c *gin.Context) {
	transactionID := c.Param("transactionId")

	if transactionID != "" {
		c.JSON(http.StatusOK, "OK, "+transactionID)
	} else {
		c.JSON(http.StatusOK, "OK")
	}
}
