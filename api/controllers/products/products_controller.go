// products_controller.go

package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}
