package router

import (
	v1 "marketplace-api/application/modules/Transaction/controllers/v1"
	v2 "marketplace-api/application/modules/Transaction/controllers/v2"

	"github.com/gin-gonic/gin"
)

//TransactionRoutes: Routes for Transaction api's v1
func TransactionRoutesV1(ctr *v1.TransactionController, r *gin.RouterGroup) {
	transactionRoutes := r.Group("/transaction")

	transactionRoutes.POST("/create", ctr.Create)

	transactionRoutes.GET("/", ctr.Read)
	transactionRoutes.GET("/:transactionID", ctr.Read)

	transactionRoutes.PUT("/update", ctr.Update)

	transactionRoutes.DELETE("/delete", ctr.Delete)

	transactionRoutes.PATCH("/pay", ctr.Pay)

	transactionRoutes.PATCH("/checkout", ctr.Checkout)
}

//TransactionRoutes: Routes for Transaction api's v2
func TransactionRoutesV2(ctr *v2.TransactionController, r *gin.RouterGroup) {
	transactionRoutes := r.Group("/transaction")

	transactionRoutes.POST("/create", ctr.Create)

	transactionRoutes.GET("/", ctr.Read)
	transactionRoutes.GET("/:transactionID", ctr.Read)

	transactionRoutes.PUT("/update", ctr.Update)

	transactionRoutes.DELETE("/delete", ctr.Delete)

	transactionRoutes.PATCH("/pay", ctr.Pay)

	transactionRoutes.PATCH("/checkout", ctr.Checkout)
}
