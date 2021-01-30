//TODO

package base

import "github.com/gin-gonic/gin"

//BaseController: Base controler interface
type BaseController interface {
	//Create: Receive HTTP request to insert data
	Create(c *gin.Context)

	//Read: Receive HTTP request to read filtered data
	Read(c *gin.Context)

	//Update: Receive HTTP request to update selected data
	Update(c *gin.Context)

	//Delete: Receive HTTP request to delete data
	Delete(c *gin.Context)
}
