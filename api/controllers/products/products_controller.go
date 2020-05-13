// products_controller.go

package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ono5/money-boy/api/domain/products"
	"github.com/ono5/money-boy/api/utils/errors"
)

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	// newProduct, saveErr := services.CreateProduct(product)
	// if saveErr != nil {
	// 	c.JSON(saveErr.Status, saveErr)
	// 	return
	// }
	// c.JSON(http.StatusCreated, newProduct)
	c.JSON(http.StatusCreated, product)
}

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}
