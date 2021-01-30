package router

import (
	TransactionControllerV1 "marketplace-api/application/modules/Transaction/controllers/v1"
	TransactionControllerV2 "marketplace-api/application/modules/Transaction/controllers/v2"
	TransactionRouter "marketplace-api/application/modules/Transaction/router"

	"github.com/gin-gonic/gin"
)

var app = gin.Default()

// Define current version of API

//Init: Register all routes in module here
func Init(version string) *gin.Engine {
	var routes *gin.RouterGroup = app.Group(version)

	if version == "v1" {
		TransactionRouter.TransactionRoutesV1(new(TransactionControllerV1.TransactionController), routes)
	} else if version == "v2" {
		TransactionRouter.TransactionRoutesV2(new(TransactionControllerV2.TransactionController), routes)
	}

	return app
}
